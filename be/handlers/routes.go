package handlers

import (
	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")

	// UI endpoints
	r.HandleFunc("/api/ui/files", GetFilesHandler).Methods("GET")
	r.HandleFunc("/api/ui/files/{filename}", GetFileHandler).Methods("GET")

	// AI endpoints
	r.HandleFunc("/api/ai/charactersdetails", CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/charactersdetails/{filename}", CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline", TimelineHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline/{filename}", TimelineHandler).Methods("GET")

	r.PathPrefix("/").Methods("GET").HandlerFunc(ReactHandler)

	return r
}
