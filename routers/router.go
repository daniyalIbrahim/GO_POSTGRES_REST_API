package routers

import (
	"fastReturns/middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	// API ENDPOINTS FOR CONSIGNMENTS
	router.HandleFunc("/api/consignment/{id}", middleware.GetOneConsignment).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/consignment/all", middleware.AllConsignments).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/consignment/update/{id}", middleware.UpdateConsignment).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/consignment/new", middleware.PostConsignment).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/consignment/delete/{id}", middleware.DeleteConsignment).Methods("DELETE", "OPTIONS")

	//API ENDPOINTS FOR SUPPLIERS
	router.HandleFunc("/api/supplier/{id}", middleware.GetOneSupplier).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/supplier/all", middleware.GetAllSupplier).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/supplier/new", middleware.PostSupplier).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/supplier/update/{id}", middleware.UpdateSupplier).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/supplier/delete/{id}", middleware.DeleteSupplier).Methods("DELETE", "OPTIONS")

	//API ENDPOINTS FOR CUSTOMERS
	router.HandleFunc("/api/customer/new", middleware.PostCustomer).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/customer/delete/{id}", middleware.DeleteCustomer).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/customer/{id}", middleware.GetOneCustomer).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/customer/update/{id}", middleware.UpdateCustomer).Methods("PUT", "OPTIONS")

	//API ENDPOINTS FOR DELIVERY VAN
	router.HandleFunc("/api/van/new", middleware.PostVan).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/van/delete/{van_id}", middleware.DeleteVan).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/van/{id}", middleware.GetOneVan).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/van/update/{id}", middleware.UpdateVan).Methods("PUT", "OPTIONS")

	return router
}
