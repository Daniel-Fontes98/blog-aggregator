package main

import (
	"blog-aggregator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env variables: %v", err)
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Couldn't connect to db: %v", err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}


	r := chi.NewRouter()
	r.Use(middlewareCors)

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handler_readiness)
	v1Router.Get("/err", handler_error)
	v1Router.Post("/users", apiCfg.handler_users_post)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handler_users_get))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handler_feeds_post))
	v1Router.Get("/feeds", apiCfg.handler_feeds_get_all)

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsPost))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.handlerFeedFollowsDelete)
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGetAllByUser))

	r.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	fmt.Printf("Server listening on port %v...\n", port)
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}