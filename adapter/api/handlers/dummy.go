package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	// 	"github.com/javiorfo/go-microservice/adapter/api/handlers"
	"github.com/javiorfo/go-microservice/common/security"
	"github.com/javiorfo/go-microservice/common/tracing"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
)

func DummyHandler(app fiber.Router, sec security.Securizer, ds service.DummyService) {
	app.Get("/dummy/:id", sec.SecureWithRoles("CLIENT_ADMIN"), func(c *fiber.Ctx) error {
		param := c.Params("id")
		log.Infof("%s Find dummy by ID: %v", tracing.LogTraceAndSpan(c), param)

		id, err := strconv.Atoi(param)
		if err != nil {
            log.Error("Invalid ID")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		dummy, err := ds.FindById(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(DummyErrorResponse(err))
		}

		return c.JSON(DummySuccessResponse(dummy))
	})
}

func DummySuccessResponse(data *model.Dummy) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func DummyErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
