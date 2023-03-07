package users

import (
	"errors"
	"lectronic/src/databases/orm/models"

	"gorm.io/gorm"
)

type User_repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) User_repo {
	return User_repo{db}
}

func (r *User_repo) GetAllUsers() (*models.Users, error) {

	var user models.Users

	result := r.db.Select("user_id, username, email, gender, image, mobile_number").Where("role = ?", "user").Order("created_at DESC").Find(&user).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &user, nil

}

func (r *User_repo) GetByID(ID string) (*models.User, error) {

	var data models.User

	result := r.db.Select("user_id, username, email, gender, role, image, mobile_number, created_at, updated_at").Find(&data, "user_id = ?", ID).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

func (r *User_repo) Register(userData *models.User) (*models.User, error) {

	result := r.db.Create(userData)
	if result.Error != nil {
		return nil, result.Error
	}

	userData.Password = ""
	userData.MobileNumber = ""
	userData.Role = ""
	userData.Image = ""

	return userData, nil
}

func (r *User_repo) UpdateUser(userData *models.User, ID string) (*models.User, error) {

	result := r.db.Model(userData).Where("user_id = ?", ID).Updates(&userData).Find(&userData).Error
	if result != nil {
		return nil, errors.New("update failed")
	}

	userData.Password = ""
	userData.Role = ""

	return userData, nil
}

func (r User_repo) UpdateToken(id, token string) error {

	var data models.User

	err := r.db.Model(data).Where("user_id = ?", id).Update("token_verify", token).Error
	if err != nil {
		return errors.New("update failed")
	}

	return nil
}

func (r *User_repo) DeleteUser(ID string) error {

	var data models.User

	result := r.db.Delete(data, "user_id = ?", ID).Error
	if result != nil {
		return result
	}

	return nil
}

func (r *User_repo) EmailExists(email string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *User_repo) UsernameExists(userName string) (bool, error) {

	var count int64

	err := r.db.Model(&models.User{}).Where("username = ?", userName).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *User_repo) GetEmail(email string) (*models.User, error) {

	var data models.User

	result := r.db.First(&data, "email = ?", email).Error
	if result != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

func (r *User_repo) GetByToken(token string) (*models.User, error) {

	var data models.User

	err := r.db.First(&data, "token_verify = ?", token).Error
	if err != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}

func (r *User_repo) TokenExists(token string) bool {

	var data models.User

	err := r.db.First(&data, "token_verify = ?", token).Error

	return err == nil
}
