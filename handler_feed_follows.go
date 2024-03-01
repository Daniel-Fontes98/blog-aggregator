package main

import (
	"blog-aggregator/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowsPost (w http.ResponseWriter, r *http.Request, u database.User) {
	type reqParameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &reqParameters{}
	err := decoder.Decode(params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID: uuid.New(),
		FeedID: params.FeedId,
		UserID: u.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Printf("Error: %v", err.Error())
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed follow")
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerFeedFollowsDelete (w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "feedFollowID")
	feedFollowId, err := uuid.Parse(param)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = cfg.DB.DeleteFeedFollowsById(r.Context(), feedFollowId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Id not found")
		return
	}

	respondWithJSON(w, http.StatusNoContent, "")
}

func (cfg *apiConfig) handlerFeedFollowsGetAllByUser (w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollows, err := cfg.DB.GetAllFeedFollowsFromUser(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get feed follows")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}