package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raad-dego/rssagg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DBURL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("dbURL must be set.")

	}
	dbQueries := database.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	v1Router := chi.NewRouter()
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	corsMux := middlewareCors(router)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
