package mocks

import (
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/internal/pagination"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockDummyService struct {
	mock.Mock
}

func (m *MockDummyService) FindById(id uint) (*model.Dummy, error) {
	args := m.Called(id)
	if dummy, ok := args.Get(0).(*model.Dummy); ok {
		return dummy, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyService) FindAll(page pagination.Page) ([]model.Dummy, error) {
	args := m.Called(page)
	if dummies, ok := args.Get(0).([]model.Dummy); ok {
		return dummies, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyService) Create(dummy *model.Dummy) error {
	args := m.Called(dummy)
	return args.Error(0)
}

// Mock Repository
type MockDummyRepository struct {
	mock.Mock
}

func (m *MockDummyRepository) FindById(id uint) (*model.Dummy, error) {
	args := m.Called(id)
	if dummy, ok := args.Get(0).(*model.Dummy); ok {
		return dummy, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyRepository) FindAll(page pagination.Page) ([]model.Dummy, error) {
    args := m.Called(page)
	if dummies, ok := args.Get(0).([]model.Dummy); ok {
		return dummies, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyRepository) Create(dummy *model.Dummy) error {
	args := m.Called(dummy)
	return args.Error(0)
}
