package router

import (
	"github.com/gorilla/mux"
	"github.com/rushikeshkandekar/controller"
	"log"
	"net/http"
)

func Router() *mux.Router {

	log.Println("serving on 9090")

	route := mux.NewRouter()

	route.HandleFunc("/customers", controller.GetAllCustomers).Methods("GET")
	route.HandleFunc("/customers", controller.CreateCustomer).Methods("POST")
	route.HandleFunc("/customers/{id}", controller.GetCustomerById).Methods("GET")
	http.ListenAndServe(":9090", route)

	return route
}
