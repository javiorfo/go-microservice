package service

import (
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/repository"
)

type DummyService interface {
	FindById(id int) (*model.Dummy, error)
	FindAll() ([]*model.Dummy, error)
	Create(model.Dummy) error
}

type dummyService struct {
	repository repository.DummyRepository
}

// func NewService(r Repository) Service {
func NewDummyService(r repository.DummyRepository) DummyService {
	return &dummyService{
		repository: r,
	}
}

func (service *dummyService) FindById(id int) (*model.Dummy, error) {
	return service.repository.FindById(id)
}

func (service *dummyService) FindAll() ([]*model.Dummy, error) {
	return nil, nil
}

func (service *dummyService) Create(d model.Dummy) error {
	return service.repository.Create(d)
}
