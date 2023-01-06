package handlers

import (
	"github.com/gorilla/mux"
	"github.com/space-devops/api-mountebank/pkg/builder"
	"github.com/space-devops/api-mountebank/pkg/client"
	"github.com/space-devops/api-mountebank/pkg/config"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if client.IsGetMethod(r) {
		client.AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := client.ExtractCID(r)
	wr := builder.BuildApiResponse(http.StatusOK,
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
	getHandler(w, r, false)
}

func GetPlanetHandler(w http.ResponseWriter, r *http.Request) {
	getHandler(w, r, true)
}

func GetSecrets(w http.ResponseWriter, r *http.Request) {
	if client.IsGetMethod(r) {
		client.AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := client.ExtractCID(r)
	wr := builder.BuildApiResponse(http.StatusOK,
		config.GetSecrets(),
		cid)

	defer func() {
		logger.LogInfo("GetSecrets Handler finished successfully", cid, logger.LogExtraInfo{
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

func getHandler(w http.ResponseWriter, r *http.Request, pathVariable bool) {
	if client.IsGetMethod(r) {
		client.AddStatusCode(&w, http.StatusMethodNotAllowed)
		return
	}

	cid := client.ExtractCID(r)
	surl := client.GetServiceURL("list")
	if pathVariable {
		planet := mux.Vars(r)["planet"]
		surl = client.GetServiceURL(planet)
	}

	bodyBytes, err := client.CallService(http.MethodGet, surl, cid)
	if err != nil {
		http.Error(w, "Error while calling external service", http.StatusInternalServerError)
		return
	}

	client.LogResponse(w, bodyBytes, cid, pathVariable)

	createResponse(&w, bodyBytes)
}

func createResponse(w *http.ResponseWriter, body []byte) {
	client.AddHeader(w, "Content-Type", "application/json")
	client.AddStatusCode(w, http.StatusOK)
	client.AddBody(w, body)
}
