package controller

import (
	"github.com/pawoo-dev/pawoo-be/dao"
	"github.com/pawoo-dev/pawoo-be/dto"
	third_party "github.com/pawoo-dev/pawoo-be/third_party/google"
)

type CompanyController interface {
	CreateCompany(companyName string, placeId string) (dto.Company, error)
	GetAllCompany() ([]dto.Company, error)
	GetCompanyDetails(companyId int) (dto.Company, error)
}

type CompanyControllerImpl struct {
}

func (c *CompanyControllerImpl) CreateCompany(companyName string, placeId string) (dto.Company, error) {
	googleClient := third_party.NewGoogleClient()
	details, err := googleClient.GetPlaceDetails(placeId)
	if err != nil {
		return dto.Company{}, err
	}
	return dao.Db.CreateCompany(dto.Company{
		Name:     companyName,
		PlaceId:  placeId,
		Location: details.FormattedAddress,
	})
}

func (c *CompanyControllerImpl) GetAllCompany() ([]dto.Company, error) {
	return dao.Db.GetAllCompany()
}

func (c *CompanyControllerImpl) GetCompanyDetails(companyId int) (dto.Company, error) {
	company, err := dao.Db.GetCompany(companyId)
	if err != nil {
		return dto.Company{}, err
	}
	serviceList, err := dao.Db.GetServiceByCompany(companyId)
	if err != nil {
		return dto.Company{}, err
	}
	company.Service = serviceList
	return company, nil
}

var CompanyControllerObj CompanyController

func NewCompanyController() {
	CompanyControllerObj = &CompanyControllerImpl{}
}
