package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mjrdev/internal/config"
	"github.com/mjrdev/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		log.Fatal("erro ao carregar .env")
	}

	router.Router(r)
	config.Db()
	config.Rdb()

	fmt.Println("listening on port 3000")
	http.ListenAndServe(":3000", r)
}
