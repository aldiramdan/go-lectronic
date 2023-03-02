package routers

import (
	"lectronic/src/databases/orm"
	"lectronic/src/modules/v1/products"
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

	return mainRoute, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("letronic backend golang!"))
}
