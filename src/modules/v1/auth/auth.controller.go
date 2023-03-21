package auth

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"net/http"

	"github.com/gorilla/mux"
)

type auth_ctrl struct {
	svc auth_service
}

func NewAuthCTRL(svc auth_service) *auth_ctrl {
	return &auth_ctrl{svc}
}

func (c *auth_ctrl) Login(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	c.svc.Login(&user).Send(w)
}

func (c *auth_ctrl) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	
	token, ok := vars["token"]

	if !ok {
		libs.GetResponse("Token not found", 404, true).Send(w)
		return
	}

	c.svc.VerifyEmail(token).Send(w)
}

func (c *auth_ctrl) ResendEmail(w http.ResponseWriter, r *http.Request) {
	
	var data models.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		libs.GetResponse(err.Error(), 400, true)
		return
	}

	c.svc.ResendEmail(&data).Send(w)
}