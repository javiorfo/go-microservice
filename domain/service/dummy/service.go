package dummy

import (
	"github.com/javiorfo/go-microservice/adapter/persistence/dummy"
	"github.com/javiorfo/go-microservice/domain/model"
)

type Service interface {
    FindById(id int) (*model.Dummy, error)
}

type service struct{
    repository dummy.Repository
}

// func NewService(r Repository) Service {
func NewService(r dummy.Repository) Service {
	return &service{
		repository: r,
	}
}

func (d *service) FindById(id int) (*model.Dummy, error) {
	return &model.Dummy{Info: "info"}, nil
}
