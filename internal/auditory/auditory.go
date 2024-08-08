package auditory

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Auditable struct {
	CreatedBy      string     `json:"-"`
	LastModifiedBy *string    `json:"-"`
	CreateDate     time.Time  `json:"-" gorm:"autoCreateTime"`
	LastModified   *time.Time `json:"-" gorm:"autoUpdateTime"`
}

func GetTokenUser(c *fiber.Ctx) string {
	if tokenUser := c.Locals("tokenUser"); tokenUser != nil {
		return tokenUser.(string)
	}
	return "unknown"
}
