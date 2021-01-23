package routes

import (
	//...
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/pgxpool"
	"github.com/rs/cors"
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
	cors := Cors()
	router.Use(cors.Handler)

	router.Route("/api", func(rt chi.Router) {
		rt.Mount("/scores", Scores(pool))
	})

	return router
}

func Cors() *cors.Cors {

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8083", "http://localhost:8081"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return cors
}
