package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/space-devops/api-mountebank/cmd/grpc"
	"github.com/space-devops/api-mountebank/pkg/config"
	"github.com/space-devops/api-mountebank/pkg/handlers"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/middleware"
	"github.com/space-devops/api-mountebank/pkg/utils"
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
		Addr:         fmt.Sprintf(":%d", cfg.Server.Http.Port),
		WriteTimeout: utils.IntToSeconds(cfg.Server.Http.WriteTimeoutSeconds),
		ReadTimeout:  utils.IntToSeconds(cfg.Server.Http.ReadTimeoutSeconds),
	}

	msg := fmt.Sprintf("Mountebank adapter listening on port %d", cfg.Server.Http.Port)
	logger.LogDebug(msg, utils.NoCorrelationId)

	go grpc.StartGRPC()

	if err := srv.ListenAndServe(); err != nil {
		logger.LogPanic(err.Error(), utils.NoCorrelationId)
	}
}
