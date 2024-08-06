package dao

import (
	"github.com/pawoo-dev/pawoo-be/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	CreateUser(user dto.User) (dto.User, error)
	GetUserByEmail(userEmail string) (dto.User, error)
}

var (
	Db Database
)

type dbImpl struct {
	Dsn          string
	DbController *gorm.DB
}

func InitDB(dsn string) error {
	if dbObj, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	} else {
		Db = NewDatabase(dsn, dbObj)
		return nil
	}
}

func NewDatabase(dsn string, dbObj *gorm.DB) Database {
	return &dbImpl{
		Dsn:          dsn,
		DbController: dbObj,
	}
}
