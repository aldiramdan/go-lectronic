package auth

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"lectronic/src/modules/v1/users"
	"os"
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

	if !user.IsActive {
		return libs.GetResponse("You account is not verified", 401, true)
	}

	jwt := libs.NewToken(user.UserID, user.Role)

	token, err := jwt.CreateToken()
	if err != nil {
		libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(tokenRes{Token: token}, 200, false)
}

func (s *auth_service) VerifyEmail(token string) *libs.Response {
	
	tokenExists := s.repo.TokenExists(token)
	if !tokenExists {
		return libs.GetResponse("verification failed", 401, true)
	}

	user, err := s.repo.GetByToken(token)
	if err != nil {
		return libs.GetResponse("user does not exist", 401, true)
	}
	
	if user.IsActive {
		return libs.GetResponse("The account has been registered", 401, true)
	}

	var data models.User

	data.IsActive = true

	_, err = s.repo.UpdateUser(&data, user.UserID)
	if err != nil {
		return libs.GetResponse("User does not exist", 401, true)
	}

	res := map[string]string{"message": "Email has been verified"}

	return libs.GetResponse(res, 200, false)
}

func (s *auth_service) ResendEmail(data *models.User) *libs.Response {
	
	emailExists, err := s.repo.EmailExists(data.Email)
	if err != nil {
		return libs.GetResponse(err.Error(), 401, true)
	}
	if emailExists {
		return libs.GetResponse("Email is not registered", 401, true)
	}

	user, err := s.repo.GetEmail(data.Email)
	if err != nil {
		return libs.GetResponse("User does not exist", 401, true)
	}

	tokenVerify, err := libs.CodeCrypt(32)
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	data.TokenVerify = tokenVerify

	emailData := libs.EmailData{
		URL: os.Getenv("BASE_URL") + "/auth/confirm_email" + tokenVerify,
		Username: data.Username,
		Subject: "Your verification code",
	}

	err = libs.SendEmail(data, &emailData)
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	err = s.repo.UpdateToken(user.UserID, tokenVerify)
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	res := map[string]string{"message": "Email is resent"}

	return libs.GetResponse(res, 200, false)
}