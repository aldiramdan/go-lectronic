package products

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/products").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/{id}", controller.GetByID).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")
	route.HandleFunc("/{id}", controller.Update).Methods("PUT")
	route.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
}
