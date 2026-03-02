package handlers

import (
	"fmt"
	"net/http"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: implement the endpoint to return a random payload for testing purposes
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "❤️")
}
