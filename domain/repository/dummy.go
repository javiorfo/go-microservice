package repository

import (
	"errors"
	"fmt"

	"github.com/javiorfo/go-microservice/domain/model"
	"gorm.io/gorm"
)

type DummyRepository interface {
	FindById(id int) (*model.Dummy, error)
	FindAll() ([]*model.Dummy, error)
	Create(model.Dummy) error
}

type dummyRepository struct {
	*gorm.DB
}

func NewDummyRepository(db *gorm.DB) DummyRepository {
	return &dummyRepository{db}
}

func (repository *dummyRepository) FindById(id int) (*model.Dummy, error) {
	var dummy model.Dummy

	result := repository.Find(&dummy, "id = ?", id)

	if result.RowsAffected == 0 {
		return nil, errors.New("Dummy Not found")
	}

	return &dummy, nil
}

func (repository *dummyRepository) FindAll() ([]*model.Dummy, error) {
	return nil, nil
}

func (repository *dummyRepository) Create(d model.Dummy) error {
	result := repository.DB.Create(&d)

	if err := result.Error; err != nil {
		return fmt.Errorf("Error creating dummy %v", err)
	}
	return nil
}
