package main

import (
	"blog-aggregator/internal/auth"
	"blog-aggregator/internal/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handler_users_post(w http.ResponseWriter, r *http.Request) {
	type reqParameters struct {
		Name string `json:"name"`
	}
	
	decoder := json.NewDecoder(r.Body)
	params := &reqParameters{}
	err := decoder.Decode(params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (cfg *apiConfig) handler_users_get(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
	}

	user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}