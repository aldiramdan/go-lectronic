package routers

import (
	"lectronic/src/databases/orm"
	"lectronic/src/modules/v1/carts"
	"lectronic/src/modules/v1/products"
	"lectronic/src/modules/v1/reviews"
	"lectronic/src/modules/v1/users"
	"lectronic/src/modules/v1/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterApp() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	db, err := orm.ConnectDB()

	if err != nil {
		return nil, err
	}

	subRouter := mainRoute.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/", homeHandler).Methods("GET")

	products.New(subRouter, db)
	carts.New(subRouter, db)
	reviews.New(subRouter, db)
	users.New(subRouter, db)
	auth.New(subRouter, db)


	return mainRoute, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("letronic backend golang!"))
}
