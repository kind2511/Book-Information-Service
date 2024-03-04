package handlers

import (
	"assignment-1/utilities"
	"encoding/json"
	"net/http"
    "strings"
)

func BookCountHanlder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleBookcountGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+http.MethodGet+
			" is supported.", http.StatusNotImplemented)

		return
	}
}

func handleBookcountGetRequest(w http.ResponseWriter, r *http.Request) {
    // Get the countrycodes from url
    languageParams := r.URL.Query().Get("language")

    // separate the countrycodes
    countryCodes := strings.Split(languageParams, ",")

    // Check if any language codes are provided
    if len(countryCodes) == 0 || countryCodes[0] == "" {
        http.Error(w, "Language parameter is required", http.StatusBadRequest)
        return
    }

    // Initialize a slice to store information for each country
    var countriesBookInfo []utilities.Bookinfo

    // Iterate over each country code
    for _, countryCode := range countryCodes {
        books, err := utilities.GetBookInformation(w, countryCode)
        if err != nil {
            http.Error(w, "Not able to retrieve information about the number of books per langugae. Check spelling of entered countrycode", http.StatusNotFound)
            return
        }

        // Get total book count in Gutendex
        totalBookCount, err := utilities.GetTotalBookCount(w)
        if err != nil {
            http.Error(w, "Not able to retrieve information about the total number of books.", http.StatusNotFound)
            return
        }

        // Get all unique authors
        authors, err := utilities.GetAllAuthors(w, countryCode)
        if err != nil {
            http.Error(w, "Not able to retrieve information about the number of authors", http.StatusNotFound)
            return
        }

        // Calculate fractions
        fraction := float64(books.Count) / float64(totalBookCount.TotalCount)

        // Prepare information for the current country
        bookinfo := utilities.Bookinfo{
            Language:  countryCode,
            Books:     books.Count,
            Authors:   len(authors),
            Fraction:  fraction,
        }

        // Append information for the current country to the slice
        countriesBookInfo = append(countriesBookInfo, bookinfo)
    }

    // Set content-type to be json
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    // Encode response
    if err := json.NewEncoder(w).Encode(countriesBookInfo); err != nil {
        http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
        return
    }
}
