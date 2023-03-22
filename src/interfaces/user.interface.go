package interfaces

import (
	"lectronic/src/databases/orm/models"
	"lectronic/src/libs"
)

type UserRepo interface {
	GetAllUsers() (*models.Users, error)
	GetByID(ID string) (*models.User, error)
	Register(userData *models.User) (*models.User, error)
	UpdateUser(userData *models.User, ID string) (*models.User, error)
	UpdateToken(id, token string) error
	DeleteUser(ID string) error
	EmailExists(email string) (bool, error)
	UsernameExists(userName string) (bool, error)
	GetEmail(email string) (*models.User, error)
	GetByToken(token string) (*models.User, error)
	TokenExists(token string) bool
}

type UserSrvc interface {
	GetAllUsers() *libs.Response
	GetByID(ID string) *libs.Response
	Register(userReg *models.User) *libs.Response
	UpdateUser(userData *models.User, ID string) *libs.Response
	DeleteUser(ID string) *libs.Response
}
