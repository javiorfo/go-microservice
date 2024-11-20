package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice-lib/auditory"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/response"
	"github.com/javiorfo/go-microservice-lib/response/codes"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice/api/request"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
)

//	@Summary		Find a dummy by ID
//	@Description	Get dummy details by ID
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Dummy ID"
//	@Success		200	{object}	model.Dummy
//	@Failure		400	{object}	response.restResponseError	"Invalid ID"
//	@Failure		404	{object}	response.restResponseError	"Internal Error"
//	@Router			/dummy/{id} [get]
//	@Security		OAuth2Password
func GetDummyById(ds service.DummyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		param := c.Params("id")
		log.Infof("%s Find dummy by ID: %v", tracing.LogTraceAndSpan(c), param)

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Error("Invalid ID")
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_FIND_ERROR, "Invalid ID"))
		}

		if dummy, err := ds.FindById(uint(id)); err != nil {
			return c.Status(http.StatusNotFound).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_FIND_ERROR, err.Error()))
		} else {
			return c.JSON(dummy)
		}
	}
}

//	@Summary		List all dummies
//	@Description	Get a list of dummies with pagination
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int												false	"Page number"
//	@Param			size		query		int												false	"Size per page"
//	@Param			sortBy		query		string											false	"Sort by field"
//	@Param			sortOrder	query		string											false	"Sort order (asc or desc)"
//	@Success		200			{object}	response.RestResponsePagination[model.Dummy]	"Paginated list of dummies"
//	@Failure		400			{object}	response.restResponseError						"Invalid query parameters"
//	@Failure		500			{object}	response.restResponseError						"Internal server error"
//	@Router			/dummy [get]
//	@Security		OAuth2Password
func GetDummies(ds service.DummyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := c.Query("page", "1")
		s := c.Query("size", "10")
		sb := c.Query("sortBy", "id")
		so := c.Query("sortOrder", "asc")

		log.Infof("%s Listing dummies...", tracing.LogTraceAndSpan(c))
		log.Infof("%s page %s, size %s, sortBy %s, sortOrder %s ", tracing.LogTraceAndSpan(c), p, s, sb, so)

		page, err := pagination.ValidateAndGetPage(p, s, sb, so)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_FIND_ERROR, err.Error()))
		}

		dummies, err := ds.FindAll(*page)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.JSON(response.RestResponsePagination[model.Dummy]{
			Pagination: pagination.Paginator(*page, len(dummies)),
			Elements:   dummies,
		})
	}
}

//	@Summary		Create a new dummy item
//	@Description	Create a new dummy item with the provided information
//	@Tags			dummy
//	@Accept			json
//	@Produce		json
//	@Param			dummy	body		request.Dummy	true	"Dummy information"
//	@Success		201		{object}	model.Dummy
//	@Failure		400		{object}	response.restResponseError	"Invalid request body or validation errors"
//	@Failure		500		{object}	response.restResponseError	"Internal server error"
//	@Router			/dummy [post]
//	@Security		OAuth2Password
func CreateDummy(ds service.DummyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dummyRequest := new(request.Dummy)

		if err := c.BodyParser(dummyRequest); err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_CREATE_ERROR, "Invalid request body"))
		}
		validate := validator.New()
		if err := validate.Struct(dummyRequest); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return c.Status(fiber.StatusBadRequest).
				JSON(response.NewRestResponseErrorWithCodeAndMsg(c, codes.DUMMY_CREATE_ERROR, validationErrors.Error()))
		}

		log.Infof("%s Received dummy: %+v", tracing.LogTraceAndSpan(c), dummyRequest)

		dummy := model.Dummy{
			Info: dummyRequest.Info,
			Auditable: auditory.Auditable{
				CreatedBy: security.GetTokenUsername(c),
			},
		}
		err := ds.Create(&dummy)

		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(response.InternalServerError(c, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(dummy)
	}
}
