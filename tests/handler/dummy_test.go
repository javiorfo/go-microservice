package dummy_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/handlers"
	"github.com/javiorfo/go-microservice/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDummyHandler(t *testing.T) {
	app := fiber.New()
	mockSec := new(mocks.MockSecurizer)
	mockService := new(mocks.MockDummyService)

	// Register the handler
	handlers.DummyHandler(app, mockSec, mockService)

	// Test cases
	tests := []struct {
		id           string
		mockReturn   interface{}
		mockError    error
		expectedCode int
	}{
		{"1", "dummyData", nil, http.StatusOK},
		{"invalid", nil, nil, http.StatusBadRequest},
		{"2", nil, errors.New("not found"), http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			// Set up the mock expectations
			if tt.id != "invalid" {
				mockService.On("FindById", mock.AnythingOfType("uint")).Return(tt.mockReturn, tt.mockError)
			}

			// Create a request
			req := httptest.NewRequest("GET", "/dummy/"+tt.id, nil)
			resp, err := app.Test(req)

			// Assert the response code
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			// If the request was valid, assert the response body
			if tt.expectedCode == http.StatusOK {
				var responseBody map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, "dummyData", responseBody["dummy"])
			}

			// Assert that the expectations were met
			mockService.AssertExpectations(t)
		})
	}
}
