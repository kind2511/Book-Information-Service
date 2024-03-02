package handlers

import (
	"assignment-1/utilities"
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
)

func ReadershipHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleReadershipGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			" is supported.", http.StatusNotImplemented)

		return
	}
}

func handleReadershipGetRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the country codes from the URL
	parts := strings.Split(r.URL.Path, "/")

	// Get the countrycode
	countryCode := parts[4]

	// Gets name and countrycodes of all countries speaking the countryCodes langugage 
	infoList, err := utilities.GetCountryNameAndCode(w, countryCode)
	if err != nil {
		return
	}

	// Parse query parameter to get the limit
	query := r.URL.Query()
	limitStr := query.Get("limit")
	var limit int
	// if limit is not empty
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	} else {
		// If limit parameter is not provided, the limit is same as num countries provided from infoList
		limit = len(infoList)
	}

	// Slice of information of each country
	var countriesInfo []utilities.CountryInfo

	// Goes through each countrycode returned from 
	for i, info := range infoList {

		// stop if equal to limit
		if i >= limit {
			break
		}

		// Get bookcount
		books, err := utilities.GetBookInformation(w, info.Isocode)
        if err != nil {
            return
        }
		
		// Get unique authors
		authors, err := utilities.GetAllAuthors(w, info.Isocode)
        if err != nil {
            return
        }

		// Get a countries population
		populations, err := utilities.GetReadership(w, info.Isocode)
		if err != nil {
			return
		}

		// Create and populate struct
		countryInfo := utilities.CountryInfo{
			Country:    info.Country,
			Isocode:    info.Isocode,
			Books:      books.Count,
			Authors:    len(authors),
			Readership: populations[0].Readership,
		}
		countriesInfo = append(countriesInfo, countryInfo)
	}

	// set content-type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode response
	if err := json.NewEncoder(w).Encode(countriesInfo); err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
