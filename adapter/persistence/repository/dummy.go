package repository

import (
	"github.com/javiorfo/go-microservice/adapter/persistence"
	"github.com/javiorfo/go-microservice/adapter/persistence/mappers"
	"github.com/javiorfo/go-microservice/domain/model"
	"gorm.io/gorm"
)

type DummyRepository interface {
    FindById(id int) (*model.Dummy, error)
}

type dummyRepository struct{
    *gorm.DB
}

func NewDummyRepository(db *gorm.DB) DummyRepository {
	return &dummyRepository{db}
}

func (repository *dummyRepository) FindById(id int) (*model.Dummy, error) {
    var entity persistence.DummyEntity

    repository.Find(&entity, "id = ?", id)

    dummy := mappers.DummyEntityToDummy(entity)

	return &dummy, nil
}
