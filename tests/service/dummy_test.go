package dummy_test

import (
	"context"
	"errors"
	"testing"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	mockClient := new(mocks.MockClient)
	mockAsync := new(mocks.MockAsync)
	dummyService := service.NewDummyService(mockRepo, mockClient, mockAsync)

	id := uint(1)
	expectedDummy := &model.Dummy{ID: id}

	ctx := context.Background()
	mockRepo.On("FindById", ctx, id).Return(expectedDummy, nil)

	result, err := dummyService.FindById(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, expectedDummy, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByIdNotFound(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	mockClient := new(mocks.MockClient)
	mockAsync := new(mocks.MockAsync)
	dummyService := service.NewDummyService(mockRepo, mockClient, mockAsync)

	id := uint(1)

	ctx := context.Background()
	mockRepo.On("FindById", ctx, id).Return(nil, errors.New("not found"))

	result, err := dummyService.FindById(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	mockClient := new(mocks.MockClient)
	mockAsync := new(mocks.MockAsync)
	dummyService := service.NewDummyService(mockRepo, mockClient, mockAsync)

	page := pagination.Page{Page: 1, Size: 10, SortBy: "id", SortOrder: "asc"}
	expectedDummies := []model.Dummy{
		{ID: 1},
		{ID: 2},
	}

	ctx := context.Background()
	mockRepo.On("FindAll", ctx, page).Return(expectedDummies, nil)

	result, err := dummyService.FindAll(ctx, page)

	assert.NoError(t, err)
	assert.Equal(t, expectedDummies, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	mockClient := new(mocks.MockClient)
	mockAsync := new(mocks.MockAsync)
	dummyService := service.NewDummyService(mockRepo, mockClient, mockAsync)

	newDummy := &model.Dummy{ID: 1}

	ctx := context.Background()
	mockRepo.On("Create", ctx, newDummy).Return(nil)

	err := dummyService.Create(ctx, newDummy)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
