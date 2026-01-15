package main

import (
	"fmt"
	"net/http"
	"time"

	"writing/config"
	"writing/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	env := config.GetEnv()

	log.WithField("port", env.Port).Info("Server starting")
	srv := &http.Server{
		Handler: handlers.Setup(),
		Addr:    fmt.Sprintf(":%d", env.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("Server failed to start")
	}
}
