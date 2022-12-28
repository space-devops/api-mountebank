package handlers

import (
	"fmt"
	"github.com/space-devops/api-mountebank/pkg/config"
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/objects"
	"github.com/space-devops/api-mountebank/pkg/utils"
	"io"
	"net/http"
)

func IsGetMethod(req *http.Request) bool {
	return req.Method != http.MethodGet
}

func ExtractCID(req *http.Request) string {
	return req.Context().Value(config.GetConfig().Global.CorrelationIdHeader).(string)
}

func AddHeader(w *http.ResponseWriter, key string, value string) {
	(*w).Header().Set(key, value)
}

func AddStatusCode(w *http.ResponseWriter, statusCode int) {
	(*w).WriteHeader(statusCode)
}

func AddBody(w *http.ResponseWriter, body []byte) {
	(*w).Write(body)
}

func GetServiceURL(name string) string {
	maps := config.GetConfig().Mountebank.Imposters
	var rt string = ""

	for _, mp := range maps {
		if mp.Name == name {
			host := config.GetConfig().Mountebank.Host

			rt = buildServiceURL(host, mp.Port, mp.Path)
		}
	}

	return rt
}

func buildServiceURL(host string, port int, path string) string {
	return fmt.Sprintf("http://%s:%d/%s", host, port, path)
}

func CallService(method string, url string, req *http.Request) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.LogPanic(
			fmt.Sprintf("Cannot create request to %s: %v", url, err),
			ExtractCID(req),
		)
		return nil, err
	}

	req.Header.Add("Accept", utils.DefaultContentType)
	req.Header.Add("Content-Type", utils.DefaultContentType)

	resp, err := client.Do(req)
	if err != nil {
		logger.LogPanic(
			fmt.Sprintf("Cannot connect to %s: %v", url, err),
			ExtractCID(req),
		)
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.LogPanic(
			fmt.Sprintf("Cannot read body from %s: %v", url, err),
			ExtractCID(req),
		)
		return nil, err
	}

	return bodyBytes, nil
}

func LogResponse(w http.ResponseWriter, bodyBytes []byte, cid string, pathVariable bool) {
	if pathVariable {
		var planet objects.PlanetWrapper
		if err := utils.JsonObjectToObject(bodyBytes, &planet, cid); err != nil {
			http.Error(w, "Error unmarshalling responses", http.StatusInternalServerError)
			return
		}

		func() {
			logger.LogInfo("GetHandlers finished successfully", cid, logger.LogExtraInfo{
				Key:   "Response",
				Value: planet,
			})
		}()
	} else {
		var planetList objects.PlanetList
		if err := utils.JsonObjectToObject(bodyBytes, &planetList, cid); err != nil {
			http.Error(w, "Error unmarshalling responses", http.StatusInternalServerError)
			return
		}

		func() {
			logger.LogInfo("GetHandlers finished successfully", cid, logger.LogExtraInfo{
				Key:   "Response",
				Value: planetList,
			})
		}()
	}
}
