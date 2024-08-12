package mocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

// Mock Security
type MockSecurizer struct {
	mock.Mock
}

func (m *MockSecurizer) SecureWithRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
