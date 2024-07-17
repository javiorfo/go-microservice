package handlers

import (
	"log"
	"net/http"

	"go.opentelemetry.io/otel"

	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/application/in"
// 	"github.com/javiorfo/go-microservice/common/security"
	"github.com/javiorfo/go-microservice/domain/model"
)

var tracer = otel.Tracer("go-microservice/dummy")

func FindById(service in.FindByIdUseCase[*model.Dummy]) fiber.Handler {
	return func(c *fiber.Ctx) error {
        _, span := tracer.Start(c.Context(), "Find Dummy by ID")
        log.Printf("TraceID: %v SpanID: %v", span.SpanContext().TraceID(), span.SpanContext().SpanID())
/*         span.SpanContext().SpanID()
        span.SpanContext().TraceID() */
	    defer span.End()
        // TODO security
/*         if err := security.SecureEndpoint(c); err != nil {
            return c.JSON(BookErrorResponse(err))
        } */

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
