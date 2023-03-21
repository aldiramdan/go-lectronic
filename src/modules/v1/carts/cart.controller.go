package carts

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/interfaces"
	"lectronic/src/libs"
	"net/http"

	"github.com/gorilla/mux"
)

type CartController struct {
	srvc interfaces.CartSrvc
}

func NewController(srvc interfaces.CartSrvc) *CartController {
	return &CartController{srvc}
}

func (c *CartController) GetAll(w http.ResponseWriter, r *http.Request) {
	c.srvc.GetAll().Send(w)
}

func (c *CartController) GetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.GetByID(id).Send(w)
}

func (c *CartController) GetByUserID(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	c.srvc.GetByUserID(user_id.(string)).Send(w)
}

func (c *CartController) Add(w http.ResponseWriter, r *http.Request) {

	user_id := r.Context().Value("user")

	var data *models.Cart

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	data.UserID = user_id.(string)

	c.srvc.Add(data).Send(w)
}

func (c *CartController) Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	var data *models.Cart

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.srvc.Update(id, data).Send(w)
}

func (c *CartController) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok {
		libs.GetResponse("ID is required", 400, true).Send(w)
		return
	}

	c.srvc.Delete(id).Send(w)
}
