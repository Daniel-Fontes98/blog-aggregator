package main

import "net/http"

func handler_error(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}