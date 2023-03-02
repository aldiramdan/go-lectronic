package auth

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"lectronic/src/modules/v1/users"
)

type auth_service struct {
	repo users.User_repo
}

type tokenRes struct {
	Token string `json:"token"`
}

func NewAuthService(repo users.User_repo) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *models.User) *libs.Response {
	
	user, err := s.repo.GetEmail(body.Email)
	if err != nil {
		return libs.GetResponse("Email or Password is incorrect", 401, true)
	}

	if libs.CheckPassword(body.Password, user.Password) {
		return libs.GetResponse("Email or Password is incorrect", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)

	token, err := jwt.CreateToken()
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(tokenRes{Token: token}, 200, false)
}