package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rfermann/gdq-stats-backend/internal/config"
	"github.com/rfermann/gdq-stats-backend/internal/services"
)

type application struct {
	db       *sqlx.DB
	config   *config.Config
	services *services.Services
}

func main() {
	config := config.New()

	db, err := sqlx.Open("pgx", config.Database_Url)
	if err != nil {
		panic(err)
	}

	app := &application{
		db:       db,
		config:   config,
		services: services.New(db),
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            true,
	}))

	router.Use(middleware.Compress(5))
	router = app.initGraphQL(router)

	fmt.Printf("starting server at port %d\n", app.config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), router)
	if err != nil {
		panic(err)
	}
}
