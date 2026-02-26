package handlers

import (
	"fmt"
	"net/http"
	"writing/ai"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func CharacterDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling character details request")

	vars := mux.Vars(r)
	filename := vars["filename"]

	prompt := ai.CharacterDetailsPrompt
	if filename != "" {
		log.WithField("filename", filename).Debug("Received request for character details with file")
		prompt = fmt.Sprintf(ai.CharacterDetailsPromptWithFile, filename)
	}

	response, err := ai.Chat(
		ai.SystemPrompt,
		prompt,
		ai.CharacterDetailsSchema,
	)

	if err != nil {
		log.WithFields(map[string]any{
			"error": err,
		}).Error("failed to get character details")
		http.Error(w, "Failed to get character details", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}
