package dummy

import (
	"github.com/javiorfo/go-microservice/domain/model"
)

type Service interface {
    FindById(id int) (*model.Dummy, error)
}

type service struct{
// repository
}

// func NewService(r Repository) Service {
func NewService() Service {
	return &service{
// 		repository: r,
	}
}

func (d *service) FindById(id int) (*model.Dummy, error) {
	return &model.Dummy{Info: "info"}, nil
}
