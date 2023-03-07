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
	route.HandleFunc("/search/", controller.Search).Methods("GET")
	route.HandleFunc("/popular/", controller.Popular).Methods("GET")
	route.HandleFunc("/", middleware.Handler(controller.Add, middleware.AuthCloudUploadFile(), middleware.AuthMiddle("admin"))).Methods("POST")
	route.HandleFunc("/{id}", middleware.Handler(controller.Update, middleware.AuthCloudUploadFile(), middleware.AuthMiddle("admin"))).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Handler(controller.Delete, middleware.AuthMiddle("admin"))).Methods("DELETE")
}
