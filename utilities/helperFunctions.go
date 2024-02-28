package utilities

import (
	"encoding/json"
	"net/http"
)

// Gets information about languages, authors, and number of books
func GetBookInformation(w http.ResponseWriter, languageCode string) (BookInfoTemp, error)  {
	// empty list of books
	var emptyBookInfo BookInfoTemp

	url := "http://129.241.150.113:8000/books/?languages=" + languageCode

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

	var books BookInfoTemp

	// Decoding JSON
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&books); err != nil {
		http.Error(w, "Did not manage to decode: "+err.Error(), http.StatusInternalServerError)
		return emptyBookInfo, err
	}

	return books, nil
}

func GetAllAuthors(w http.ResponseWriter, languageCode string) ([]Authors, error) {
    uniqueAuthors := make(map[string]bool)
    var allAuthors []Authors

    nextPage := "http://129.241.150.113:8000/books/?languages=" + languageCode
    for nextPage != "" {
        r, err := http.NewRequest(http.MethodGet, nextPage, nil)
        if err != nil {
            http.Error(w, "Error in creating request", http.StatusInternalServerError)
            return nil, err
        }
        r.Header.Add("content-type", "application/json")

        client := &http.Client{}
        defer client.CloseIdleConnections()

        res, err := client.Do(r)
        if err != nil {
            http.Error(w, "Did not manage to issue request", http.StatusInternalServerError)
            return nil, err
        }

        var pageResponse struct {
            Next     string    `json:"next"`
            Previous string    `json:"previous"`
            Results  []Authors `json:"results"`
        }

        decoder := json.NewDecoder(res.Body)
        if err := decoder.Decode(&pageResponse); err != nil {
            http.Error(w, "Did not manage to decode: "+err.Error(), http.StatusInternalServerError)
            return nil, err
        }

        for _, authors := range pageResponse.Results {
            for _, author := range authors.Authors {
                if !uniqueAuthors[author.Name] {
                    uniqueAuthors[author.Name] = true
                    allAuthors = append(allAuthors, authors)
                }
            }
        }

        nextPage = pageResponse.Next
    }

    return allAuthors, nil
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

