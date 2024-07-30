package dummy

import (
	"github.com/javiorfo/go-microservice/adapter/persistence"
	"github.com/javiorfo/go-microservice/domain/model"
)

type Repository interface {
    FindById(id int) (*model.Dummy, error)
}

type repository struct{
// repository
}

// func NewService(r Repository) Service {
func NewRepository() Repository {
	return &repository{
// 		repository: r,
	}
}

func (d *repository) FindById(id int) (*model.Dummy, error) {
    dummy := persistence.DummyEntityToDummy(persistence.DummyEntity{Id: 1, Info: "info"})
	return &dummy, nil
}
