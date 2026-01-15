package handlers

import (
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/api/files", GetFilesHandler).Methods("GET")
	r.HandleFunc("/api/files/{filename}", GetFileHandler).Methods("GET")

	r.PathPrefix("/").Methods("GET").HandlerFunc(ReactHandler)

	return r
}
