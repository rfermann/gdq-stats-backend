package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rfermann/gdq-stats-backend/internal/config"
	"github.com/rfermann/gdq-stats-backend/internal/services"
)

type application struct {
	db       *sql.DB
	config   *config.Config
	services *services.Services
}

func main() {
	config := config.New()

	db, err := sql.Open("pgx", config.Database_Url)
	if err != nil {
		panic(err)
	}

	app := &application{
		db:       db,
		config:   config,
		services: services.New(db),
	}
	router := chi.NewRouter()
	router = app.initGraphQL(router)

	fmt.Printf("starting server at port %d\n", app.config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), router)
	if err != nil {
		panic(err)
	}
}
