package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/chaosystema/go-microservice/adapter/in/routes/handlers"
	"github.com/chaosystema/go-microservice/application/in"
	"github.com/chaosystema/go-microservice/domain/model"
)

func DummyRouter(app fiber.Router, service in.FindByIdUseCase[*model.Dummy]) {
// 	app.Get("/chaosystema/dummy", SecureEndpoint)
	app.Get("/dummy", handlers.FindById(service))
}
