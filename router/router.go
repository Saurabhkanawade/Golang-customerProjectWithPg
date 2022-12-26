package router

import (
	"github.com/Saurabhkanawade/golang-practice/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//router.HandleFunc("/api/customers", controller.CreateCustomer).Methods("POST", "GET")
	router.HandleFunc("/api/customers", controller.GetCustomers).Methods("POST")
	router.HandleFunc("/api/customers", controller.GetCustomers).Methods("GET")
	router.HandleFunc("/api/customers/{id}", controller.GetCustomerById).Methods("GET")
	router.HandleFunc("/api/customers/{id}", controller.UpdateCustomerById).Methods("PUT")
	router.HandleFunc("/api/customers/{id}", controller.DeleteCustomerById).Methods("DELETE")

	return router
}
