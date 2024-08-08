package dao

import "github.com/pawoo-dev/pawoo-be/dto"

func (db *dbImpl) CreateCompany(company dto.Company) (dto.Company, error) {
	result := db.DbController.Create(&company)
	return company, result.Error
}

func (db *dbImpl) GetAllCompany() ([]dto.Company, error) {
	var companyList []dto.Company

	result := db.DbController.Find(&companyList)
	return companyList, result.Error
}

func (db *dbImpl) GetCompany(companyName string) (dto.Company, error) {
	var company dto.Company
	results := db.DbController.Where("name = ?", companyName).First(&company)
	return company, results.Error
}
