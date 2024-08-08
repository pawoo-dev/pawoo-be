package controller

import (
	"github.com/pawoo-dev/pawoo-be/dao"
	"github.com/pawoo-dev/pawoo-be/dto"
)

type CompanyController interface {
	CreateCompany(companyName string) (dto.Company, error)
	GetAllCompany() ([]dto.Company, error)
}

type CompanyControllerImpl struct {
}

func (c *CompanyControllerImpl) CreateCompany(companyName string) (dto.Company, error) {
	return dao.Db.CreateCompany(dto.Company{
		Name: companyName,
	})
}

func (c *CompanyControllerImpl) GetAllCompany() ([]dto.Company, error) {
	return dao.Db.GetAllCompany()
}

var CompanyControllerObj CompanyController

func NewCompanyController() {
	CompanyControllerObj = &CompanyControllerImpl{}
}
