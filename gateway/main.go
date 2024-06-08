package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	common "github.com/huxleyberg/omsv2-commons"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // Adds some basic logging

	h := NewHandler()
	h.registerRoutes(r)

	log.Printf("starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, r); err != nil {
		log.Fatalf("failed to start http server - +%v", err)
	}
}
