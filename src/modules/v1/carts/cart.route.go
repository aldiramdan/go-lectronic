package carts

import (
	"lectronic/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/carts").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", middleware.Handler(controller.GetAll, middleware.AuthMiddle("admin"))).Methods("GET")
	route.HandleFunc("/id={id}", middleware.Handler(controller.GetByID, middleware.AuthMiddle("admin", "user"))).Methods("GET")
	route.HandleFunc("/carts-list", middleware.Handler(controller.GetByUserID, middleware.AuthMiddle("user"))).Methods("GET")
	route.HandleFunc("/", middleware.Handler(controller.Add, middleware.AuthMiddle("user"))).Methods("POST")
	route.HandleFunc("/{id}", middleware.Handler(controller.Update, middleware.AuthMiddle("user"))).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Handler(controller.Delete, middleware.AuthMiddle("user"))).Methods("DELETE")
}
