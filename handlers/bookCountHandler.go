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

	parts := strings.Split(r.URL.String(), "/")

	books, err := utilities.GetBookInformation(w, parts[4])
	if err != nil {
		return
	}

	totalBookCount, err := utilities.GetTotalBookCountfunc(w)
	if err != nil {
		return
	}

	// Get the unique authors
	uniqueAuthors := make(map[string]bool)
    for _, book := range books.Results {
        for _, author := range book.Authors {
            uniqueAuthors[author.Name] = true
        }
    }

	BooksInfo := utilities.Bookinfo {
		Language: parts[4],
		Books: books.Count,
		Authors: len(uniqueAuthors),
		Fraction: float64(books.Count) / float64(totalBookCount.TotalCount),
	}
	
	w.Header().Add("content-type", "apllication-json")

	encoder := json.NewEncoder(w)
	
	err = encoder.Encode(BooksInfo)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
