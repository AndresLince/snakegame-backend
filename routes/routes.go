package routes

import (
	//...
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/pgxpool"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func Routers(pool *pgxpool.Pool) *chi.Mux {
	router := chi.NewRouter()

	//router.Use(middleware.Logger)

	router.Route("/api", func(rt chi.Router) {
		rt.Mount("/scores", Scores(pool))
	})

	return router
}
