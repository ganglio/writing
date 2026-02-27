package ai

import (
	"fmt"
	"net/http"
	"writing/ai"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func TimelineHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling timeline request")

	vars := mux.Vars(r)
	filename := vars["filename"]

	prompt := ai.StoryTimelinePrompt
	if filename != "" {
		log.WithField("filename", filename).Debug("Received request for story timeline with file")
		prompt = fmt.Sprintf(ai.StoryTimelinePromptWithFile, filename)
	}

	response, err := ai.Chat(
		ai.SystemPrompt,
		prompt,
		ai.StoryTimelineSchema,
	)

	if err != nil {
		log.WithFields(map[string]any{
			"error": err,
		}).Error("failed to get story timeline")
		http.Error(w, "Failed to get story timeline", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}
