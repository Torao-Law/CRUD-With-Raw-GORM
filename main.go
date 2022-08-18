package main

import (
	"be-waybucks/database"
	"be-waybucks/pkg/mysql"
	"be-waybucks/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// handler main
func main() {
	// initial database
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
