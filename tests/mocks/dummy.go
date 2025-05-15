package mocks

import (
	"context"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockDummyService struct {
	mock.Mock
}

func (m *MockDummyService) FindById(ctx context.Context, id uint) (*model.Dummy, error) {
	args := m.Called(ctx, id)
	if dummy, ok := args.Get(0).(*model.Dummy); ok {
		return dummy, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyService) FindAll(ctx context.Context, page pagination.Page, info string) ([]model.Dummy, error) {
	args := m.Called(ctx, page, info)
	if dummies, ok := args.Get(0).([]model.Dummy); ok {
		return dummies, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyService) Create(ctx context.Context, dummy *model.Dummy) error {
	args := m.Called(ctx, dummy)
	return args.Error(0)
}

func (m *MockDummyService) External(ctx context.Context) (string, error) {
	args := m.Called(ctx)
	if res, ok := args.Get(0).(string); ok {
		return res, args.Error(1)
	}
	return "", args.Error(1)
}

// Mock Repository
type MockDummyRepository struct {
	mock.Mock
}

func (m *MockDummyRepository) FindById(ctx context.Context, id uint) (*model.Dummy, error) {
	args := m.Called(ctx, id)
	if dummy, ok := args.Get(0).(*model.Dummy); ok {
		return dummy, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyRepository) FindAll(ctx context.Context, page pagination.Page, info string) ([]model.Dummy, error) {
	args := m.Called(ctx, page, info)
	if dummies, ok := args.Get(0).([]model.Dummy); ok {
		return dummies, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockDummyRepository) Create(ctx context.Context, dummy *model.Dummy) error {
	args := m.Called(ctx, dummy)
	return args.Error(0)
}
