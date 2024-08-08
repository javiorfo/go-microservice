package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice/common/response"
	"github.com/javiorfo/go-microservice/common/security"
	"github.com/javiorfo/go-microservice/common/tracing"
	"github.com/javiorfo/go-microservice/domain/service"
)

func DummyHandler(app fiber.Router, sec security.Securizer, ds service.DummyService) {
    // Get by ID
	app.Get("/dummy/:id", sec.SecureWithRoles("CLIENT_ADMIN"), func(c *fiber.Ctx) error {
		param := c.Params("id")
		log.Infof("%s Find dummy by ID: %v", tracing.LogTraceAndSpan(c), param)

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Error("Invalid ID")
            invalidParameter := response.NewRestResponseError(response.ResponseError{
				Code:    "DUMMY_ERROR",
				Message: "Invalid ID",
			})
			return c.Status(fiber.StatusBadRequest).JSON(invalidParameter)
		}

		dummy, err := ds.FindById(id)
		if err != nil {
			dummyNotFound := response.NewRestResponseError(response.ResponseError{
				Code:    "DUMMY_ERROR",
				Message: err.Error(),
			})
			return c.Status(http.StatusNotFound).JSON(dummyNotFound)
		}

		return c.JSON(fiber.Map{"dummy": dummy})
	})
    
    // List with pagination and sorting
    app.Get("/dummy", sec.SecureWithRoles("CLIENT_ADMIN"), func(c *fiber.Ctx) error {
		log.Infof("%s List dummies", tracing.LogTraceAndSpan(c))

		dummies, err := ds.FindAll()
		if err != nil {
			dummyNotFound := response.NewRestResponseError(response.ResponseError{
				Code:    "DUMMY_ERROR",
				Message: err.Error(),
			})
			return c.Status(http.StatusNotFound).JSON(dummyNotFound)
		}

		return c.JSON(fiber.Map{"dummy": dummies})
	})

    // Create
}
