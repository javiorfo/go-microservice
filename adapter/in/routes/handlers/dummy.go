package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/application/in"
	"github.com/javiorfo/go-microservice/common/security"
	"github.com/javiorfo/go-microservice/domain/model"
)

func FindById(service in.FindByIdUseCase[*model.Dummy]) fiber.Handler {
	return func(c *fiber.Ctx) error {
        // TODO security
        if err := security.SecureEndpoint(c); err != nil {
            return err
        }

		fetched, err := service.FindById(1)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(BookErrorResponse(err))
		}
		return c.JSON(BookSuccessResponse(fetched))
	}
}

func BookSuccessResponse(data *model.Dummy) *fiber.Map {
	book := model.Dummy{
		Info: data.Info,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
