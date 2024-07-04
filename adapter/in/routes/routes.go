package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/adapter/in/routes/handlers"
	"github.com/javiorfo/go-microservice/application/in"
	"github.com/javiorfo/go-microservice/domain/model"
)

func DummyRouter(app fiber.Router, service in.FindByIdUseCase[*model.Dummy]) {
// 	app.Get("/app/dummy", SecureEndpoint)
	app.Get("/dummy", handlers.FindById(service))
}
