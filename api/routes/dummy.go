package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/handlers"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice-lib/security"
)

const keycloakRoles = "CLIENT_ADMIN"
const root = "/dummy"

func Dummy(app fiber.Router, sec security.Securizer, service service.DummyService) {
	app.Get(root+"/:id", sec.SecureWithRoles(keycloakRoles), handlers.GetDummyById(service))
	app.Get(root, sec.SecureWithRoles(keycloakRoles), handlers.GetDummies(service))
	app.Post(root, sec.SecureWithRoles(keycloakRoles), handlers.CreateDummy(service))
}
