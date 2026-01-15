package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"writing/config"
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

	http.Error(w, "Not yet implemented", http.StatusInternalServerError)
}
