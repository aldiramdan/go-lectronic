package users

import (
	"encoding/json"
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

type user_ctrl struct {
	svc *user_service
}

func NewUserCTRL(svc *user_service) *user_ctrl {
	return &user_ctrl{svc}
}

func (c *user_ctrl) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	c.svc.GetAllUsers().Send(w)
}

func (c *user_ctrl) GetByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	user_id := r.Context().Value("user")

	c.svc.GetByID(user_id.(string)).Send(w)
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

	user_id := r.Context().Value("user").(string)

	var user models.User

	imageName := r.Context().Value("imageName").(string)
	user.Image = imageName

	err := schema.NewDecoder().Decode(&user, r.MultipartForm.Value)
	if err != nil {
		log.Printf("%v", err)
		libs.GetResponse(err.Error(), 400, true).Send(w)
		return
	}

	c.svc.UpdateUser(&user, user_id).Send(w)
}

func (c *user_ctrl) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	user_id := r.Context().Value("user")

	c.svc.DeleteUser(user_id.(string)).Send(w)
}
