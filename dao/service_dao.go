package dao

import (
	"fmt"

	"github.com/pawoo-dev/pawoo-be/dto"
)

func (db *dbImpl) CreateService(service dto.Service) (dto.Service, error) {
	result := db.DbController.Create(&service)
	return service, result.Error
}

func (db *dbImpl) GetServiceByCompany(companyId int) ([]dto.Service, error) {
	// Get all service
	var serviceList []dto.Service

	result := db.DbController.Where("company_id = ?", companyId).Find(&serviceList)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(serviceList) == 0 {
		return nil, fmt.Errorf("result not found")
	}
	// for each service get all options
	for index, s := range serviceList {
		var optionList []dto.ServiceOptions
		result := db.DbController.Where("service_id = ?", s.ID).Find(&optionList)
		if result.Error != nil {
			return nil, result.Error
		}
		serviceList[index].Options = optionList
	}
	// return service
	return serviceList, nil
}

func (db *dbImpl) AddOptions(serviceId int, serviceOption []dto.ServiceOptions) ([]dto.ServiceOptions, error) {
	for index, _ := range serviceOption {
		serviceOption[index].ServiceId = serviceId
	}

	result := db.DbController.Create(&serviceOption)

	return serviceOption, result.Error
}

func (db *dbImpl) UpdateOptions(serviceOption dto.ServiceOptions) error {
	results := db.DbController.Updates(serviceOption)
	if results.RowsAffected == 0 {
		return fmt.Errorf("no entry found please refresh page")
	}
	return results.Error
}

func (db *dbImpl) UpdateService(service dto.Service) error {
	results := db.DbController.Updates(service)
	if results.RowsAffected == 0 {
		return fmt.Errorf("no entry found please refresh page")
	}
	return results.Error
}
