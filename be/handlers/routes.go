package handlers

import (
	"writing/config"
	"writing/handlers/ai"

	"github.com/gorilla/mux"
)

var (
	env = config.GetEnv()
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")

	// UI endpoints
	r.HandleFunc("/api/ui/files", GetFilesHandler).Methods("GET")
	r.HandleFunc("/api/ui/files/{filename}", GetFileHandler).Methods("GET")

	if env.IsDev() {
		r.HandleFunc("/api/ui/random", RandomHandler).Methods("POST")
	}

	// AI endpoints
	r.HandleFunc("/api/ai/charactersdetails", ai.CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/charactersdetails/{filename}", ai.CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline", ai.TimelineHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline/{filename}", ai.TimelineHandler).Methods("GET")
	r.HandleFunc("/api/ai/conversations", ai.ConversationAmountHandler).Methods("GET")
	r.HandleFunc("/api/ai/conversations/{filename}", ai.ConversationAmountHandler).Methods("GET")

	// React app endpoint
	r.PathPrefix("/").Methods("GET").HandlerFunc(ReactHandler)

	return r
}
