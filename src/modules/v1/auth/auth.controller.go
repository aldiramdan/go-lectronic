package auth

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"net/http"
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