package utilities

import (
	"encoding/json"
	"net/http"
)

// Gets information about languages, authors, and number of books
func GetBookInformation(w http.ResponseWriter, languageCode string) ([]BookInfo, error)  {
	// empty list of books
	var emptyBookInfo []BookInfo

	url := "http://129.241.150.113:8000/books/?language=" + languageCode

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error in creating request", http.StatusInternalServerError)
		return emptyBookInfo, err
	}

	// Setting content type
	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		http.Error(w, "Did not manage to issue request", http.StatusInternalServerError)
		return emptyBookInfo, err
	}

	var books []BookInfo

	// Decoding JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&books); err != nil {
		http.Error(w, "Did not manage to decode: "+err.Error(), http.StatusInternalServerError)
		return emptyBookInfo, err
	}

	// If no books found
	if len(books) == 0 {
		http.Error(w, "No books found", http.StatusNotFound)
		return emptyBookInfo, err
	}

	return books, nil
}

// Gets the total number of books in the Gutendex API
func GetTotalBookCountfunc(w http.ResponseWriter) (TotalBookCount, error)  {
	
	var emptyBookCountInfo TotalBookCount

	url := "http://129.241.150.113:8000/books/"

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error in creating request", http.StatusInternalServerError)
		return emptyBookCountInfo, err
	}

	// Setting content type
	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		http.Error(w, "Did not manage to issue request", http.StatusInternalServerError)
		return emptyBookCountInfo, err
	}

	var totalBookCount TotalBookCount

	// Decoding JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&totalBookCount); err != nil {
		http.Error(w, "Did not manage to decode: "+err.Error(), http.StatusInternalServerError)
		return totalBookCount, err
	}

	return totalBookCount, nil
}

