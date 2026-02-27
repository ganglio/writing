package ai

import (
	"fmt"
	"net/http"
	"writing/ai"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ConversationAmountHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling conversation amount request")

	vars := mux.Vars(r)
	filename := vars["filename"]

	prompt := ai.ConversationAmountPrompt
	if filename != "" {
		log.WithField("filename", filename).Debug("Received request for conversation amount with file")
		prompt = fmt.Sprintf(ai.ConversationAmountPromptWithFile, filename)
	}

	response, err := ai.Chat(
		ai.SystemPrompt,
		prompt,
		ai.ConversationAmountSchema,
	)

	if err != nil {
		log.WithFields(map[string]any{
			"error": err,
		}).Error("failed to get conversation amount")
		http.Error(w, "Failed to get conversation amount", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}
