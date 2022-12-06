package handlers

import (
	"encoding/json"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	wr := utils.BuildApiResponse(http.StatusOK,
		"Welcome to Mountebank Sidecar",
		3001)

	jr, je := json.Marshal(wr)
	if je != nil {
		http.Error(w, "Error marshalling responses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jr)
}
