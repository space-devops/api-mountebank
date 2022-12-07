package middleware

import (
	"context"
	"fmt"
	"github.com/space-devops/mountebank-sidecar/pkg/config"
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(config.GetConfig().Global.CorrelationIdHeader)
		if correlationID == "" {
			correlationID = utils.GenerateCorrelationId()
		}

		defer log(r, correlationID)

		ctx := context.WithValue(r.Context(), config.GetConfig().Global.CorrelationIdHeader, correlationID)
		w.Header().Set(config.GetConfig().Global.CorrelationIdHeader, correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func log(r *http.Request, cid string) {
	logger.LogInfo("CorrelationMiddleware completed successfully", cid,
		logger.LogExtraInfo{
			Key:   "Request - Method",
			Value: r.Method,
		}, logger.LogExtraInfo{
			Key:   "Request - Path",
			Value: fmt.Sprintf("%s %s", r.Host, r.URL.Path),
		})
}
