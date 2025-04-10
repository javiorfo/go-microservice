package mocks

import (
	"github.com/javiorfo/go-microservice-lib/integration"
	"github.com/stretchr/testify/mock"
)

// Mock Client
type MockClient struct {
	mock.Mock
}

func (m *MockClient) Send(r integration.Request) (*integration.Response[integration.RawData], error) {
	args := m.Called()
	if dummy, ok := args.Get(0).(*integration.Response[integration.RawData]); ok {
		return dummy, args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock Async
type MockAsync struct {
	mock.Mock
}

func (m *MockAsync) Execute(r integration.Request) {
	m.Called(r)
}
