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

	infoList, err := utilities.GetCountryNameAndCode(w, countryCode)
	if err != nil {
		return
	}

	// Parse query parameter to get the limit
	query := r.URL.Query()
	limitStr := query.Get("limit")
	var limit int
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
	} else {
		// If limit parameter is not provided, set limit to a large value to effectively have no limit
		limit = len(infoList)
	}

	var countriesInfo []utilities.CountryInfo

	for i, info := range infoList {

		if i >= limit {
			break
		}

		books, err := utilities.GetBookInformation(w, info.Isocode)
        if err != nil {
            return
        }
	
		authors, err := utilities.GetAllAuthors(w, info.Isocode)
        if err != nil {
            return
        }

		populations, err := utilities.GetReadership(w, info.Isocode)
		if err != nil {
			return
		}

		population := 0
        if len(populations) > 0 {
            population = populations[0].Readership
        }

		countryInfo := utilities.CountryInfo{
			Country:    info.Country,
			Isocode:    info.Isocode,
			Books:      books.Count,
			Authors:    len(authors),
			Readership: population,
		}
		countriesInfo = append(countriesInfo, countryInfo)
	}

	// Encode response
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(countriesInfo); err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

