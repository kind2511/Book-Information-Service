package handlers

import (
	"assignment-1/utilities"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This service does not provide any functionality on root path level.\n\nPlease use paths:\n"+utilities.BOOKCOUNT_PATH+"?language={:two_letter_language_code+}/, \n"+utilities.READERSHIP_PATH+"{:two_letter_language_code}{?limit={:number}}/, \n"+utilities.STATUS_PATH, http.StatusOK)
}