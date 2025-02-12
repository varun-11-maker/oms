package router

import (
	"usingPostgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/order/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/order", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/neworder", middleware.CreateStock).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
