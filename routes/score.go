package routes

import (
	"net/http"

	"../handlers"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/pgxpool"
)

func Scores(pool *pgxpool.Pool) http.Handler {
	scoresHandler := handlers.NewScoresHandler(pool)
	router := chi.NewRouter()

	router.Get("/", scoresHandler.GetAll)
	router.Post("/", scoresHandler.Create)
	return router
}
