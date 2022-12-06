package middleware

import (
	"context"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(utils.CorrelationIdHeaderName)
		if correlationID == "" {
			correlationID = utils.GenerateCorrelationId()
		}

		ctx := context.WithValue(r.Context(), utils.CorrelationIdHeaderName, correlationID)
		w.Header().Set(utils.CorrelationIdHeaderName, correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
