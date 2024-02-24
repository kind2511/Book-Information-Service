package handlers

import (
	"assignment-1/utilities"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// declared variable of type time.Time
var startTime time.Time

// Sets startTime
func init() {
    startTime = time.Now()
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	// Gets statuscode of Gutendex API
	gutenDexApiStatusCode, err := http.Get("http://129.241.150.113:8000/books/")
	if err != nil {
		err = fmt.Errorf("error occurred while making HTTP request: %v", err)
		fmt.Println(err)
		return
	}

	// Gets statuscode of Language2countries API
	languageApiStatusCode, err := http.Get("http://129.241.150.113:3000/language2countries/")
	if err != nil {
		err = fmt.Errorf("error occurred while making HTTP request: %v", err)
		fmt.Println(err)
		return
	}

	// Gets statuscode of Countries API
	countriesApiStatusCode, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		err = fmt.Errorf("error occurred while making HTTP request: %v", err)
		fmt.Println(err)
		return
	}

	/// Time the server has been up since start
    uptime := time.Since(startTime).Seconds()

	// Creates a Status struct
	Status := utilities.Status {
		Gutendexapi: gutenDexApiStatusCode.StatusCode,
		Langugaeapi: languageApiStatusCode.StatusCode,
		Countriesapi: countriesApiStatusCode.StatusCode,
		Version: "v1",
		Uptime: uptime,
	}

	// Sets the content-type to be json
	w.Header().Add("content-type", "apllication-json")

	encoder := json.NewEncoder(w)
	
	err = encoder.Encode(Status)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
