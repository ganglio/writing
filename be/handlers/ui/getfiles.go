package ui

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"writing/config"
)

func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	env := config.GetEnv()

	var files []string
	err := fs.WalkDir(os.DirFS(env.WorkFolder), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && !strings.HasPrefix(path, ".") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read files: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}
