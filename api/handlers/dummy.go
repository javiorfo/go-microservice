package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice/api/request"
	"github.com/javiorfo/go-microservice/internal/auditory"
	"github.com/javiorfo/go-microservice/internal/response"
	"github.com/javiorfo/go-microservice/internal/response/codes"
	"github.com/javiorfo/go-microservice/internal/security"
	"github.com/javiorfo/go-microservice/internal/tracing"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
)

var keycloakRoles = "CLIENT_ADMIN"

func DummyHandler(app fiber.Router, sec security.Securizer, ds service.DummyService) {
	// Get by ID
	app.Get("/dummy/:id", sec.SecureWithRoles(keycloakRoles), func(c *fiber.Ctx) error {
		param := c.Params("id")
		log.Infof("%s Find dummy by ID: %v", tracing.LogTraceAndSpan(c), param)

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Error("Invalid ID")
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(codes.DUMMY_FIND_ERROR, "Invalid ID"))
		}

		dummy, err := ds.FindById(id)
		if err != nil {
			return c.Status(http.StatusNotFound).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(codes.DUMMY_FIND_ERROR, err.Error()))
		}

		return c.JSON(fiber.Map{"dummy": dummy})
	})

	// List with pagination and sorting
	app.Get("/dummy", sec.SecureWithRoles(keycloakRoles), func(c *fiber.Ctx) error {
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
	app.Post("/dummy", sec.SecureWithRoles(keycloakRoles), func(c *fiber.Ctx) error {
		var dummyRequest request.Dummy

		if err := c.BodyParser(&dummyRequest); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(codes.DUMMY_CREATE_ERROR, "Invalid request body"))
		}

		log.Infof("%s Received dummy: %+v", tracing.LogTraceAndSpan(c), dummyRequest)

		err := ds.Create(model.Dummy{
			Info: dummyRequest.Info,
			Auditable: auditory.Auditable{
				CreatedBy: auditory.GetTokenUser(c),
			},
		})

		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(codes.DUMMY_CREATE_ERROR, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(dummyRequest)
	})
}
