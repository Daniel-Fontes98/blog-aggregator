package main

import "net/http"

func handler_readiness(w http.ResponseWriter, _ *http.Request) {
	respondWithJSON(w, http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	})
}