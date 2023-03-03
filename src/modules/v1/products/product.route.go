package products

import (
	"lectronic/src/middleware"

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
	route.HandleFunc("/", middleware.Handler(controller.Add, middleware.AuthCloudUploadFile())).Methods("POST")
	route.HandleFunc("/{id}", middleware.Handler(controller.Update, middleware.AuthCloudUploadFile())).Methods("PUT")
	route.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
}
