package routes

import (
	"be-waybucks/handlers"
	"be-waybucks/pkg/mysql"
	"be-waybucks/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("GET")
}
