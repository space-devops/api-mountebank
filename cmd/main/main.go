package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"log"
	"net/http"
	"time"
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

func main() {
	fmt.Println("Hello world from Golang")

	r := mux.NewRouter()
	r.HandleFunc("/", WelcomeHandler).GetHandler()

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", utils.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
