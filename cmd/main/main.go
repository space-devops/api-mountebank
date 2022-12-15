package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/space-devops/mountebank-sidecar/pkg/config"
	"github.com/space-devops/mountebank-sidecar/pkg/handlers"
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/middleware"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	logger.InitLogger(cfg.Logger.File)
	logger.SetLogLevel(cfg.Logger.Level)

	r := mux.NewRouter()

	r.Use(middleware.CorrelationMiddleware)

	r.HandleFunc("/", handlers.WelcomeHandler).GetHandler()
	r.HandleFunc("/planets", handlers.GetPlanetListHandler).GetHandler()
	r.HandleFunc("/planet/{planet}", handlers.GetPlanetHandler).GetHandler()

	r.HandleFunc("/live", handlers.LivenessHandler).GetHandler()
	r.HandleFunc("/ready", handlers.ReadinessHandler).GetHandler()

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		WriteTimeout: utils.IntToSeconds(cfg.Server.WriteTimeoutSeconds),
		ReadTimeout:  utils.IntToSeconds(cfg.Server.ReadTimeoutSeconds),
	}

	msg := fmt.Sprintf("Mountebank adapter listening on port %d", cfg.Server.Port)
	logger.LogDebug(msg, utils.NoCorrelationId)

	if err := srv.ListenAndServe(); err != nil {
		logger.LogPanic(err.Error(), utils.NoCorrelationId)
	}
}
