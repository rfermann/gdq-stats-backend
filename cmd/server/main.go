package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rfermann/gdq-stats-backend/internal/config"
)

type application struct {
	config *config.Config
}

func main() {
	app := &application{
		config: config.New(),
	}
	router := chi.NewRouter()

	fmt.Printf("starting server at port %d\n", app.config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), router)
	if err != nil {
		panic(err)
	}
}
