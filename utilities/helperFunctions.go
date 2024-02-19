package utilities

import (
	"net/http"
)

func makeGetRequest(w http.ResponseWriter, url string) (*http.Response, error) {
	// Send GET request to an api
	response, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "Error in creating request. Error code: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections() 

	// Issue request
	res, err := client.Do(response)
	if err != nil {
		http.Error(w, "Error in response. Error message: "+err.Error(), res.StatusCode)
		return nil, err
	}

	return res, nil
}
