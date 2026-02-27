package main

import (
	"fmt"
	"net/http"

	"writing/ai"
	"writing/config"
	"writing/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {
	env := config.GetEnv()

	if env.IsDev() {
		log.Info("Running in development mode")
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.Info("Running in production mode")
	}

	log.WithField("schema", ai.CharacterDetailsSchema).Debug("AI character details schema loaded")

	log.WithField("env", env).Info("Environment configuration loaded")
	log.WithField("port", env.Port).Info("Server starting")
	srv := &http.Server{
		Handler: handlers.Setup(),
		Addr:    fmt.Sprintf(":%d", env.Port),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).Fatal("Server failed to start")
	}
}
