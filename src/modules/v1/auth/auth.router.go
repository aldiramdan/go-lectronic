package auth

import (
	"lectronic/src/modules/v1/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Auth(route *mux.Router, db *gorm.DB) {

	router := route.PathPrefix("/user/auth").Subrouter()

	repo := users.NewUserRepo(db)
	service := NewAuthService(repo)
	ctrl := NewAuthCTRL(*service)

	router.HandleFunc("", ctrl.Login).Methods("POST")
}
