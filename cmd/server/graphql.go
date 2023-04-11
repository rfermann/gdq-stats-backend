package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rfermann/gdq-stats-backend/cmd/server/resolvers"
	"github.com/rfermann/gdq-stats-backend/internal/config"
	graph "github.com/rfermann/gdq-stats-backend/internal/gql"
)

func (app *application) initGraphQL(router *chi.Mux) *chi.Mux {
	graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &resolvers.Resolver{
					Services: app.services,
				},
			},
		),
	)

	router.Post("/gql/query", func(w http.ResponseWriter, r *http.Request) {
		graphqlHandler.ServeHTTP(w, r)
	})

	if app.config.Environment == config.DevEnv {
		router.Get("/gql/playground", func(w http.ResponseWriter, r *http.Request) {
			playground.AltairHandler("gdq-stats Playground", "/gql/query").ServeHTTP(w, r)
		})
	}

	return router
}
