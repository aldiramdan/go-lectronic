package users

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"lectronic/src/middleware"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type user_ctrl struct {
	svc *user_service
}

func NewUserCTRL(svc *user_service) *user_ctrl {
	return &user_ctrl{svc}
}

func (c *user_ctrl) Register(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.GetResponse(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.Register(&user).Send(w)
}

func (c *user_ctrl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	userID, ok := r.Context().Value(middleware.UserID("user")).(string)
	if !ok {
		libs.GetResponse("Unauthorized", 401, true).Send(w)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateUser(&user, string(userID)).Send(w)
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	id := r.Context().Value(middleware.UserID("user")).(string)

	c.svc.DeleteUser(id).Send(w)
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllUsers().Send(w)
}

func (c *user_ctrl) GetByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	id := r.Context().Value(middleware.UserID("user")).(string)

	c.svc.GetByID(id).Send(w)
}