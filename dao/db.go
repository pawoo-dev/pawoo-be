package dao

import (
	"github.com/pawoo-dev/pawoo-be/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	CreateUser(user dto.User) (dto.User, error)
	GetUserByEmail(userEmail string) (dto.User, error)

	CreateService(service dto.Service) (dto.Service, error)
	GetServiceByCompany(companyId int) ([]dto.Service, error)
	UpdateService(service dto.Service) error
	AddOptions(serviceId int, serviceOption []dto.ServiceOptions) ([]dto.ServiceOptions, error)
	UpdateOptions(serviceOption dto.ServiceOptions) error

	CreateCompany(company dto.Company) (dto.Company, error)
	GetAllCompany() ([]dto.Company, error)
	GetCompany(companyName string) (dto.Company, error)
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
