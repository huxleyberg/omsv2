package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	common "github.com/huxleyberg/omsv2commons"
	pb "github.com/huxleyberg/omsv2commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = common.EnvString("ORDER_SERVICE_ADDR", "localhost:3000")
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server : %v", err)
	}
	defer conn.Close()

	orderClient := pb.NewOrderServiceClient(conn)

	r := chi.NewRouter()
	r.Use(middleware.Logger) // Adds some basic logging

	h := NewHandler(orderClient)
	h.registerRoutes(r)

	log.Printf("starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, r); err != nil {
		log.Fatalf("failed to start http server - +%v", err)
	}
}
