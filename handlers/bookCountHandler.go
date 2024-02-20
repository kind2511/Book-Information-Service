package handlers

import (
	"net/http"
)


func BookCountHanlderr(w http.ResponseWriter, r *http.Request) {
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
	
}

