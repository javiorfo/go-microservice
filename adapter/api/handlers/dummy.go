package handlers

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service/dummy"
)

var tracer = otel.Tracer(fmt.Sprintf("%s/adapter/api/handlers/dummy", config.AppName))

func FindDummyById(service dummy.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
        _, span := tracer.Start(c.Context(), "FindDummyByID")
        log.Infof("TraceID: %v SpanID: %v", span.SpanContext().TraceID(), span.SpanContext().SpanID())
	    defer span.End()

		fetched, err := service.FindById(1)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(BookErrorResponse(err))
		}

		return c.JSON(BookSuccessResponse(fetched))
	}
}

func BookSuccessResponse(data *model.Dummy) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
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
