package handlers

import (
	"writing/config"
	"writing/handlers/ai"
	"writing/handlers/ui"

	"github.com/gorilla/mux"
)

var (
	env = config.GetEnv()
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")
	if env.IsDev() {
		r.HandleFunc("/api/random", RandomHandler).Methods("POST")
	}

	// UI endpoints
	r.HandleFunc("/api/ui/files", ui.GetFilesHandler).Methods("GET")
	r.HandleFunc("/api/ui/files/{filename}", ui.GetFileHandler).Methods("GET")

	// AI endpoints
	r.HandleFunc("/api/ai/charactersdetails", ai.CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/charactersdetails/{filename}", ai.CharacterDetailsHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline", ai.TimelineHandler).Methods("GET")
	r.HandleFunc("/api/ai/timeline/{filename}", ai.TimelineHandler).Methods("GET")
	r.HandleFunc("/api/ai/conversations", ai.ConversationAmountHandler).Methods("GET")
	r.HandleFunc("/api/ai/conversations/{filename}", ai.ConversationAmountHandler).Methods("GET")

	// React app endpoint
	r.PathPrefix("/").Methods("GET").HandlerFunc(ui.ReactHandler)

	return r
}
