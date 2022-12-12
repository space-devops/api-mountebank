package handlers

import (
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/objects"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if IsGetMethod(r) {
		AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := ExtractCID(r)
	wr := utils.BuildApiResponse(http.StatusOK,
		"Welcome to Mountebank Sidecar",
		cid)

	defer func() {
		logger.LogInfo("WelcomeHandler finished successfully", cid, logger.LogExtraInfo{
			Key:   "Response",
			Value: wr,
		})
	}()

	obj, err := utils.ObjectToJsonObject(wr, cid)
	if err != nil {
		http.Error(w, "Error marshalling responses", http.StatusInternalServerError)
		return
	}

	createResponse(&w, obj)
}

func GetPlanetListHandler(w http.ResponseWriter, r *http.Request) {
	if IsGetMethod(r) {
		AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := ExtractCID(r)
	surl := GetServiceURL("list")

	bodyBytes, err := CallService(http.MethodGet, surl, r)
	if err != nil {
		http.Error(w, "Error while calling external service", http.StatusInternalServerError)
		return
	}

	var planetList objects.PlanetList
	if err = utils.JsonObjectToObject(bodyBytes, &planetList, cid); err != nil {
		http.Error(w, "Error unmarshalling responses", http.StatusInternalServerError)
		return
	}

	func() {
		logger.LogInfo("WelcomeHandler finished successfully", cid, logger.LogExtraInfo{
			Key:   "Response",
			Value: planetList,
		})
	}()

	createResponse(&w, bodyBytes)
}

func createResponse(w *http.ResponseWriter, body []byte) {
	AddHeader(w, "Content-Type", "application/json")
	AddStatusCode(w, http.StatusOK)
	AddBody(w, body)
}
