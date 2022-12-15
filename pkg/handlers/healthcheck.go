package handlers

import (
	"fmt"
	"github.com/space-devops/mountebank-sidecar/pkg/config"
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if IsGetMethod(r) {
		AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	host := config.GetConfig().Mountebank.Host
	port := config.GetConfig().Mountebank.Health.Port
	path := config.GetConfig().Mountebank.Health.Path

	upstream := buildServiceURL(host, port, path)

	logger.LogInfo(
		fmt.Sprintf("Healthcheck upstream service: %s", upstream),
		utils.NoCorrelationId,
	)

	bodyBytes, err := CallService(http.MethodGet, upstream, r)
	if err != nil {
		http.Error(w, "Error while calling external service", http.StatusInternalServerError)
		return
	}

	createResponse(&w, bodyBytes)
}

func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	if IsGetMethod(r) {
		AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := utils.NoCorrelationId
	wr := utils.BuildApiResponse(http.StatusOK,
		"Liveness probe for kubernetes healthcheck system",
		cid)

	defer func() {
		logger.LogInfo("Liveness probe handler finished successfully", cid, logger.LogExtraInfo{
			Key:   "Response",
			Value: wr,
		})
	}()

	obj, err := utils.ObjectToJsonObject(wr, cid)
	if err != nil {
		http.Error(w, "Error marshalling responses on liveness handler", http.StatusInternalServerError)
		return
	}

	createResponse(&w, obj)
}
