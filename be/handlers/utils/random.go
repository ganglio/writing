package utils

import (
	"fmt"
	"net/http"

	lorem "github.com/drhodes/golorem"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: implement the endpoint to return a random payload for testing purposes
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, lorem.Sentence(10, 20))
}
