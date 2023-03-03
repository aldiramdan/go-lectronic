package users

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
	"os"

	"gorm.io/gorm"
)

type user_service struct {
	repo User_repo
}

func NewUserService(repo User_repo) *user_service {
	return &user_service{repo}
}

func (s *user_service) Register(userReg *models.User) *libs.Response {
	
	emailExists, err := s.repo.EmailExists(userReg.Email)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	if emailExists {
		return libs.GetResponse("Email already exists", 400, true)
	}

	userNameExists, err := s.repo.UsernameExists(userReg.Username)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	if userNameExists {
		return libs.GetResponse("Username already used", 400, true)
	}

	hashPass, err := libs.HashPassword(userReg.Password)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	userReg.Password = hashPass

	tokenVerify, err := libs.CodeCrypt(32)
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	userReg.TokenVerify = tokenVerify

	emailData := libs.EmailData{
		URL: os.Getenv("BASE_URL") + "/auth/confirm_email/" + tokenVerify,
		Username: userReg.Username,
		Subject: "Your verification code",
	}

	err = libs.SendEmail(userReg, &emailData)
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	data, err := s.repo.Register(userReg)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	result, _ := s.repo.GetByID(data.UserID)

	return libs.GetResponse(result, 200, false)
}

func (s *user_service) UpdateUser(userData *models.User, ID string) *libs.Response {
	
	var user models.User

	err := s.repo.db.Where("user_id = ?", ID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return libs.GetResponse("Data not found", 404, true)
		} else {
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	if userData.Password != "" {
		hashPass, err := libs.HashPassword(userData.Password)
		if err != nil {
			return libs.GetResponse("Password update failed", 400, true)
		}

		userData.Password = hashPass
	}

	emailExists, err := s.repo.EmailExists(userData.Email)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	if emailExists {
		return libs.GetResponse("Email already exists", 400, true)
	}

	userNameExists, err := s.repo.UsernameExists(userData.Username)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	if userNameExists {
		return libs.GetResponse("Username already used", 400, true)
	}
	
	if userData.Username == "" {
		userData.Username = user.Username
	}
	if userData.Email == "" {
		userData.Email = user.Email
	}
	if userData.Password == "" {
		userData.Password = user.Password
	}
	if userData.Gender == "" {
		userData.Username = user.Gender
	}
	if userData.Image == "" {
		userData.Image = user.Image
	}
	if userData.MobileNumber == "" {
		userData.MobileNumber = user.MobileNumber
	}

	result, err := s.repo.UpdateUser(userData, ID)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	return libs.GetResponse(result, 200, false)
}

func (s *user_service) DeleteUser(ID string) *libs.Response {

	_, err := s.repo.GetByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return libs.GetResponse("Data not found", 404, true)
		} else {
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	err = s.repo.DeleteUser(ID)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}

	result := map[string]string{"Message": "User has been deleted"}

	return libs.GetResponse(result, 200, false)
}

func (s *user_service) GetAllUsers() *libs.Response {
	
	data, err := s.repo.GetAllUsers()
	if err != nil {
		return libs.GetResponse(err.Error(), 500, true)
	}

	return libs.GetResponse(data, 200, false)
}

func (s *user_service) GetByID(ID string) *libs.Response {

	data, err := s.repo.GetByID(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return libs.GetResponse(err.Error(), 404, true)
		} else {
			return libs.GetResponse(err.Error(), 500, true)
		}
	}

	return libs.GetResponse(data, 200, false) 
}