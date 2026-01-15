package handlers

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"

	"writing/config"
	"writing/embed"
)

func ReactHandler(w http.ResponseWriter, r *http.Request) {
	if config.GetEnv().IsDev() {
		target, err := url.Parse(fmt.Sprintf("http://%s:%d", config.GetEnv().FEDevHost, config.GetEnv().FEDevPort))
		if err != nil {
			http.Error(w, "Failed to parse target URL", http.StatusInternalServerError)
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(w, r)
		return
	}

	// in production, serve the embedded React files
	rootFs, err := fs.Sub(embed.FeFS, "fe_build")
	if err != nil {
		http.Error(w, "Failed to access embedded files", http.StatusInternalServerError)
		return
	}
	fs := http.FS(rootFs)
	fileServer := http.FileServer(fs)
	fileServer.ServeHTTP(w, r)
}
