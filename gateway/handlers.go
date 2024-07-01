package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	common "github.com/huxleyberg/omsv2commons"
	pb "github.com/huxleyberg/omsv2commons/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(r chi.Router) {
	r.Route("/api/customers/{customerID}/orders", func(r chi.Router) {
		r.Post("/orders", h.HandleCreateOrder)
	})
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WirteError(w, http.StatusBadRequest, err.Error())
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})

	w.WriteHeader(http.StatusOK)
}
