package handlers

import (
	"encoding/json"
	"github.com/space-devops/mountebank-sidecar/pkg/config"
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cid := r.Context().Value(config.GetConfig().Global.CorrelationIdHeader).(string)

	wr := utils.BuildApiResponse(http.StatusOK,
		"Welcome to Mountebank Sidecar",
		cid)

	defer func() {
		logger.LogInfo("WelcomeHandler finished successfully", cid, logger.LogExtraInfo{
			Key:   "Response",
			Value: wr,
		})
	}()

	jr, je := json.Marshal(wr)
	if je != nil {
		http.Error(w, "Error marshalling responses", http.StatusInternalServerError)
		logger.LogError("Error marshalling responses", cid, logger.LogExtraInfo{
			Key:   "MarshallerError",
			Value: je.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
}
