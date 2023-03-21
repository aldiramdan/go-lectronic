package auth

import (
	"lectronic/src/modules/v1/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(route *mux.Router, db *gorm.DB) {

	router := route.PathPrefix("/auth").Subrouter()

	repo := users.NewUserRepo(db)
	service := NewAuthService(repo)
	ctrl := NewAuthCTRL(*service)

	router.HandleFunc("/login", ctrl.Login).Methods("POST")
	router.HandleFunc("/confirm_email/{token}", ctrl.VerifyEmail).Methods("GET")
	router.HandleFunc("/resend_email", ctrl.ResendEmail).Methods("POST")
}
