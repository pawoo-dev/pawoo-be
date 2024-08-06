package dao

import (
	"github.com/pawoo-dev/pawoo-be/dto"
)

func (db *dbImpl) CreateUser(user dto.User) (dto.User, error) {
	result := db.DbController.Create(user)
	return user, result.Error
}

func (db *dbImpl) GetUserByEmail(userEmail string) (dto.User, error) {
	var user dto.User
	result := db.DbController.Where("email = ?", userEmail).First(&user)
	return user, result.Error
}
