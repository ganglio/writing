package handlers

import (
	"net/http"

	"writing/ai"

	log "github.com/sirupsen/logrus"
)

func CharacterDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling character details request")

	response, err := ai.Chat(
		ai.SystemPrompt,
		ai.CharacterDetailsPrompt,
		ai.CharacterDetailsSchema,
	)

	if err != nil {
		log.WithFields(map[string]any{
			"error": err,
		}).Error("failed to get character details")
		http.Error(w, "Failed to get character details", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
