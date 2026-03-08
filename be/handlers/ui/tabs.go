package ui

import (
	"encoding/json"
	"net/http"

	"writing/config"
)

func Tabs(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(config.GetEnv().Tabs)
	if err != nil {
		http.Error(w, "Failed to marshal tabs", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
