package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(r chi.Router) {
	r.Route("/api/customers/{customerID}/orders", func(r chi.Router) {
		r.Post("/orders", h.HandleCreateOrder)
	})
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customerID")

	// Do something with customerID
	_ = customerID // Placeholder

	w.WriteHeader(http.StatusOK)
}
