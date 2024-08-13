package dummy_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/handlers"
	"github.com/javiorfo/go-microservice/api/request"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/internal/pagination"
	"github.com/javiorfo/go-microservice/internal/response"
	"github.com/javiorfo/go-microservice/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindById(t *testing.T) {
    app := fiber.New()
    mockSec := new(mocks.MockSecurizer)
    mockService := new(mocks.MockDummyService)

	handlers.DummyHandler(app, mockSec, mockService)

	// FIND BY ID
	tests := []struct {
		id           string
		mockReturn   *model.Dummy
		mockError    error
		expectedCode int
	}{
		{"1", &model.Dummy{Info: "info"}, nil, fiber.StatusOK},
		{"2", nil, errors.New("Dummy not found"), fiber.StatusNotFound},
		{"invalid", nil, nil, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			if tt.id != "invalid" {
				id, _ := strconv.Atoi(tt.id)
				mockService.On("FindById", uint(id)).Return(tt.mockReturn, tt.mockError)
			}

			req := httptest.NewRequest("GET", "/dummy/"+tt.id, nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if tt.expectedCode == http.StatusOK {
				var responseBody map[string]model.Dummy
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, "info", responseBody["dummy"].Info)
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestFindAll(t *testing.T) {
    app := fiber.New()
    mockSec := new(mocks.MockSecurizer)
    mockService := new(mocks.MockDummyService)

	handlers.DummyHandler(app, mockSec, mockService)
    // FIND ALL
	t.Run("Successful", func(t *testing.T) {
		page := pagination.Page{Page: 1, Size: 10, SortBy: "info", SortOrder: "asc"}
		mockService.On("FindAll", page).Return([]model.Dummy{{ID: 1, Info: "info"}}, nil)

		req := httptest.NewRequest("GET", "/dummy?page=1&size=10&sortBy=info&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var responseBody response.RestResponsePagination[model.Dummy]
		json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.Equal(t, 1, responseBody.Pagination.Total)
		assert.Equal(t, "info", responseBody.Elements[0].Info)

		mockService.AssertExpectations(t)
	})

	t.Run("DB Error", func(t *testing.T) {
		mockService.On("FindAll", mock.Anything).Return(nil, errors.New("data source error"))

		req := httptest.NewRequest("GET", "/dummy?page=1&size=10&sortBy=id&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("Pagination Bad Request", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/dummy?page=invalid&size=10&sortBy=id&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

func TestCreate(t *testing.T) {
    app := fiber.New()
    mockSec := new(mocks.MockSecurizer)
    mockService := new(mocks.MockDummyService)

	handlers.DummyHandler(app, mockSec, mockService)

	t.Run("Successful", func(t *testing.T) {
		dummyRequest := request.Dummy{Info: "dummy info"}
		mockService.On("Create", mock.Anything).Return(nil)

		body, _ := json.Marshal(dummyRequest)
		req := httptest.NewRequest("POST", "/dummy", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var responseBody map[string]model.Dummy
		json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.Equal(t, "dummy info", responseBody["dummy"].Info)

		mockService.AssertExpectations(t)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		dummyRequest := request.Dummy{Info: "failed"}
		mockService.On("Create", mock.Anything).Return(errors.New("create error"))

		body, _ := json.Marshal(dummyRequest)
		req := httptest.NewRequest("POST", "/dummy", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {
        body := `{ "invalid": 10 }`
		req := httptest.NewRequest("POST", "/dummy", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}
