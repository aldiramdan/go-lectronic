package routers

import (
	"lectronic/src/databases/orm"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterApp() (*mux.Router, error) {

	mainRoute := mux.NewRouter()

	_, err := orm.ConnectDB()

	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("letronic backend golang"))
	})

	return mainRoute, nil
}
