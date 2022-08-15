package routes

import (
	"be-waybucks/handlers"
	"be-waybucks/pkg/mysql"
	"be-waybucks/repositories"

	"github.com/gorilla/mux"
)

func TopingRoutes(r *mux.Router) {
	topingRepository := repositories.RepositoryToping(mysql.DB)
	h := handlers.HandlerToping(topingRepository)

	r.HandleFunc("/topings", h.FindToping).Methods("GET")
	r.HandleFunc("/toping/{id}", h.GetToping).Methods("GET")
	r.HandleFunc("/toping", h.CreateToping).Methods("POST")
	r.HandleFunc("/toping/{id}", h.DeleteToping).Methods("DELETE")
}
