package main

import (
	"blog-aggregator/internal/database"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) handlerGetPostsByUser (w http.ResponseWriter, r *http.Request, u database.User) {
	var defaultLimit = 0
	param := chi.URLParam(r, "limit")
	if param != "" {
		var err error
		defaultLimit, err = strconv.Atoi(param)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid limit")
			return
		}
	} else {
		defaultLimit = 10
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: u.ID,
		Limit: int32(defaultLimit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}