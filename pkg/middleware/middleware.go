package middleware

import (
	"context"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = utils.GenerateCorrelationId()
		}

		ctx := context.WithValue(r.Context(), "X-Correlation-ID", correlationID)
		w.Header().Set("X-Correlation-ID", correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
