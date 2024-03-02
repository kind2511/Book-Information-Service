package utilities

import (
	"encoding/json"
	"net/http"
)

// Makes an HTTP GET request and decodes the JSON response into the given decodeObject interface
func makeRequestAndDecodeJSON(w http.ResponseWriter, url string, decodeObject interface{}) error {
    r, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        http.Error(w, "Error in creating request", http.StatusInternalServerError)
        return err
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
        return err
    }

    // Decoding JSON
    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(decodeObject); err != nil {
        http.Error(w, "Did not manage to decode: "+err.Error(), http.StatusInternalServerError)
        return err
    }

    return nil
}

// Gets information about number of books
func GetBookInformation(w http.ResponseWriter, languageCode string) (BookCount, error) {
	// returns empyty BookCount if error
    var emptyBookInfo BookCount

    var books BookCount

    url := "http://129.241.150.113:8000/books/?languages=" + languageCode

	// make get request to get book count
    if err := makeRequestAndDecodeJSON(w, url, &books); err != nil {
        return emptyBookInfo, err
    }
    return books, nil
}

// Gets the total number of books in the Gutendex API
func GetTotalBookCount(w http.ResponseWriter) (TotalBookCount, error) {
	// returns empyt TotalBookCount if error
    var emptyBookCountInfo TotalBookCount

    var totalBookCount TotalBookCount

    url := "http://129.241.150.113:8000/books/"

	// makes get request to get total number of books in GutenDex API
    if err := makeRequestAndDecodeJSON(w, url, &totalBookCount); err != nil {
        return emptyBookCountInfo, err
    }
    return totalBookCount, nil
}

// Gets all authors
func GetAllAuthors(w http.ResponseWriter, languageCode string) ([]Authors, error) {
    uniqueAuthors := make(map[string]bool)

	// slice containg all unique authors
    var allAuthors []Authors

    nextPage := "http://129.241.150.113:8000/books/?languages=" + languageCode
	// Fetches pages til there are no more pages
    for nextPage != "" {

		// Response gutendex api struct
		var gutenDexResponse GutendexResponse

		// makes get request to retrieve all authors on all pages
        if err := makeRequestAndDecodeJSON(w, nextPage, &gutenDexResponse); err != nil {
            return nil, err
        }
		// Filters out all none-unique athors based on their names
        for _, authors := range gutenDexResponse.Results {
            for _, author := range authors.Authors {
                if !uniqueAuthors[author.Name] {
                    uniqueAuthors[author.Name] = true
                    allAuthors = append(allAuthors, authors)
                }
            }
        }
        nextPage = gutenDexResponse.Next
    }
    return allAuthors, nil
}

// Gets countrynames and iso-codes
func GetCountryNameAndCode(w http.ResponseWriter, languageCode string) ([]CountryNameAndCode, error) {
	// retuns empty slice if error
    var emptyCountryInfo []CountryNameAndCode

    var countryInfo []CountryNameAndCode
    url := "http://129.241.150.113:3000/language2countries/" + languageCode

	// makes get request to get country name and iso code
    if err := makeRequestAndDecodeJSON(w, url, &countryInfo); err != nil {
        return emptyCountryInfo, err
    }
    return countryInfo, nil
}

// Gets the populations
func GetReadership(w http.ResponseWriter, countryCode string) ([]CountryPopulation, error) {
	// retuns empty slice if error
    var emptyCountryPopulationInfo []CountryPopulation

    var populations []CountryPopulation

    url := "https://restcountries.com/v3.1/alpha/" + countryCode

	// makes get request to get population
    if err := makeRequestAndDecodeJSON(w ,url, &populations); err != nil {
        return emptyCountryPopulationInfo, err
    }
    return populations, nil
}
