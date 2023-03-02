package reviews

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/reviews").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/{id}", controller.GetByID).Methods("GET")
	route.HandleFunc("/product/{id}", controller.GetByProductID).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")
}
