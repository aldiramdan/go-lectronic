package users

import (
	"lectronic/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {
	router := route.PathPrefix("/users").Subrouter()

	repo := NewUserRepo(db)
	service := NewUserService(repo)
	ctrl := NewUserCTRL(service)

	router.HandleFunc("/", middleware.Handler(ctrl.GetAllUsers, middleware.AuthMiddle("admin"))).Methods("GET")
	router.HandleFunc("/profile", middleware.Handler(ctrl.GetByID, middleware.AuthMiddle("admin", "user"))).Methods("GET")
	router.HandleFunc("/register", ctrl.Register).Methods("POST")
	router.HandleFunc("/profile/edit", middleware.Handler(ctrl.UpdateUser, middleware.AuthCloudUploadFile(), middleware.AuthMiddle("admin", "user"))).Methods("PUT")
	router.HandleFunc("/delete", middleware.Handler(ctrl.DeleteUser, middleware.AuthMiddle("admin", "user"))).Methods("DELETE")
}
