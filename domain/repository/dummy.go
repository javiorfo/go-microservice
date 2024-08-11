package repository

import (
	"errors"
	"fmt"

	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/internal/pagination"
	"gorm.io/gorm"
)

type DummyRepository interface {
	FindById(id uint) (*model.Dummy, error)
	FindAll(pagination.Page) ([]model.Dummy, error)
	Create(*model.Dummy) error
}

type dummyRepository struct {
	*gorm.DB
}

func NewDummyRepository(db *gorm.DB) DummyRepository {
	return &dummyRepository{db}
}

func (repository *dummyRepository) FindById(id uint) (*model.Dummy, error) {
	var dummy model.Dummy

	result := repository.Find(&dummy, "id = ?", id)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Dummy Not found")
	}

	return &dummy, nil
}

func (repository *dummyRepository) FindAll(page pagination.Page) ([]model.Dummy, error) {
	var dummies []model.Dummy
	offset := (page.Page - 1)
	results := repository.Offset(offset).Limit(page.Size).Order(fmt.Sprintf("%s %s", page.SortBy, page.SortOrder)).Find(&dummies)

	if err := results.Error; err != nil {
		return nil, err
	}

	return dummies, nil
}

func (repository *dummyRepository) Create(d *model.Dummy) error {
	result := repository.DB.Create(d)

	if err := result.Error; err != nil {
		return fmt.Errorf("Error creating dummy %v", err)
	}
	return nil
}
