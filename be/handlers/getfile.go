package handlers

import (
	"fmt"
	"net/http"
	"os"
	"writing/config"

	"github.com/gorilla/mux"
)

func GetFileHandler(w http.ResponseWriter, r *http.Request) {
	env := config.GetEnv()
	vars := mux.Vars(r)
	filename := vars["filename"]

	filePath := fmt.Sprintf("%s/%s", env.WorkFolder, filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/markdown")
	w.Write(data)
}
