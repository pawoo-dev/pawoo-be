package controller

import (
	"github.com/pawoo-dev/pawoo-be/dao"
	"github.com/pawoo-dev/pawoo-be/dto"
)

type UserController interface {
	GetUserByEmail(email string) (dto.User, error)
	CreateUser(user dto.User) (dto.User, error)
}

type UserControllerImpl struct {
}

func (u *UserControllerImpl) GetUserByEmail(email string) (dto.User, error) {
	return dao.Db.GetUserByEmail(email)
}

func (u *UserControllerImpl) CreateUser(user dto.User) (dto.User, error) {
	return dao.Db.CreateUser(user)
}

var UserControllerObj UserController

func NewUserController() {
	UserControllerObj = &UserControllerImpl{}
}
