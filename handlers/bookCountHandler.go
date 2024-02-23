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
    // Parse the country codes from the URL
    parts := strings.Split(r.URL.Path, "/")
    
	// Get the countrycodes
    countryCodes := strings.Split(parts[4], ",")

    // Initialize a slice to store information for each country
    var countriesInfo []utilities.Bookinfo

    // Iterate over each country code
    for _, countryCode := range countryCodes {
        books, err := utilities.GetBookInformation(w, countryCode)
        if err != nil {
            return
        }

        // Get total book count for all countries
        totalBookCount, err := utilities.GetTotalBookCountfunc(w)
        if err != nil {
            return
        }

        // Calculate unique authors
        uniqueAuthors := make(map[string]bool)
        for _, book := range books.Results {
            for _, author := range book.Authors {
                uniqueAuthors[author.Name] = true
            }
        }

        // Calculate fraction
        fraction := float64(books.Count) / float64(totalBookCount.TotalCount)

        // Prepare information for the current country
        countryInfo := utilities.Bookinfo{
            Language:  countryCode,
            Books:     books.Count,
            Authors:   len(uniqueAuthors),
            Fraction:  fraction,
        }

        // Append information for the current country to the slice
        countriesInfo = append(countriesInfo, countryInfo)
    }

    // Encode response
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(countriesInfo); err != nil {
        http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
        return
    }
}
