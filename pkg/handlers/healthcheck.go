package handlers

import (
	"fmt"
	"github.com/space-devops/api-mountebank/pkg/builder"
	"github.com/space-devops/api-mountebank/pkg/client"
	"github.com/space-devops/api-mountebank/pkg/config"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	if client.IsGetMethod(r) {
		client.AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	host := config.GetConfig().Mountebank.Host
	port := config.GetConfig().Mountebank.Health.Port
	path := config.GetConfig().Mountebank.Health.Path

	upstream := client.BuildServiceURL(host, port, path)

	logger.LogInfo(
		fmt.Sprintf("Healthcheck upstream service: %s", upstream),
		utils.NoCorrelationId,
	)

	bodyBytes, err := client.CallService(http.MethodGet, upstream, client.ExtractCID(r))
	if err != nil {
		http.Error(w, "Error while calling external service", http.StatusInternalServerError)
		return
	}

	createResponse(&w, bodyBytes)
}

func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	if client.IsGetMethod(r) {
		client.AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := utils.NoCorrelationId
	wr := builder.BuildApiResponse(http.StatusOK,
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
