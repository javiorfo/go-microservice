package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/adapter/in/routes/handlers"
// 	"github.com/javiorfo/go-microservice/application/in"
	"github.com/javiorfo/go-microservice/common/security"
// 	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service/dummy"
)

func DummyRouter(app fiber.Router, securizer security.Securizer, service dummy.Service) {
	app.Get("/dummy", securizer.SecureWithRoles("CLIENT_ADMIN"), handlers.FindById(service))
}
