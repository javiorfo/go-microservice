package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice/api/request"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/internal/auditory"
	"github.com/javiorfo/go-microservice/internal/pagination"
	"github.com/javiorfo/go-microservice/internal/response"
	"github.com/javiorfo/go-microservice/internal/response/codes"
	"github.com/javiorfo/go-microservice/internal/security"
	"github.com/javiorfo/go-microservice/internal/tracing"
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
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_FIND_ERROR, "Invalid ID"))
		}

		if dummy, err := ds.FindById(id); err != nil {
			return c.Status(http.StatusNotFound).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_FIND_ERROR, err.Error()))
		} else {
			return c.JSON(fiber.Map{"dummy": dummy})
		}
	})

	// List with pagination and sorting
	app.Get("/dummy", sec.SecureWithRoles(keycloakRoles), func(c *fiber.Ctx) error {
		p := c.Query("page", "1")
		s := c.Query("size", "10")
		sb := c.Query("sortBy", "id")
		so := c.Query("sortOrder", "asc")

		log.Infof("%s Listing dummies...", tracing.LogTraceAndSpan(c))
		log.Infof("%s page %s, size %s, sortBy %s, sortOrder %s ", tracing.LogTraceAndSpan(c), p, s, sb, so)

		page, err := pagination.ValidateAndGetPage(p, s, sb, so)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.InternalServerError(c, err.Error()))
		}

		dummies, err := ds.FindAll(*page)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.JSON(response.RestResponsePagination[model.Dummy]{
			Pagination: pagination.Paginator(*page, len(dummies)),
			Elements:   dummies,
		})
	})

	// Create
	app.Post("/dummy", sec.SecureWithRoles(keycloakRoles), func(c *fiber.Ctx) error {
		var dummyRequest request.Dummy

		if err := c.BodyParser(&dummyRequest); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_CREATE_ERROR, "Invalid request body"))
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
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(dummyRequest)
	})
}
