package controller

import (
	"github.com/pawoo-dev/pawoo-be/dao"
	"github.com/pawoo-dev/pawoo-be/dto"
)

type ServiceController interface {
	CreateService(service dto.Service) error
	GetService(companyId int) ([]dto.Service, error)
	UpdateService(service dto.Service) error
	UpdateOptions(options dto.ServiceOptions) error
	AddOptions(serviceId int, options []dto.ServiceOptions) error
}

type ServiceControllerImpl struct {
}

var ServiceControllerObj ServiceController

func NewServiceController() {
	ServiceControllerObj = &ServiceControllerImpl{}
}

func (s *ServiceControllerImpl) CreateService(service dto.Service) error {
	service, err := dao.Db.CreateService(service)
	if err != nil {
		return err
	}

	_, err = dao.Db.AddOptions(service.ID, service.Options)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceControllerImpl) GetService(companyId int) ([]dto.Service, error) {
	return dao.Db.GetServiceByCompany(companyId)
}

func (s *ServiceControllerImpl) UpdateOptions(options dto.ServiceOptions) error {
	return dao.Db.UpdateOptions(options)
}

func (s *ServiceControllerImpl) UpdateService(service dto.Service) error {
	return dao.Db.UpdateService(service)
}

func (s *ServiceControllerImpl) AddOptions(serviceId int, options []dto.ServiceOptions) error {
	_, err := dao.Db.AddOptions(serviceId, options)
	return err
}
