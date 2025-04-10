package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice/api/handlers"
	"github.com/javiorfo/go-microservice/domain/service"
)

const keycloakRoles = "CLIENT_ADMIN"
const root = "/dummy"

func Dummy(app fiber.Router, sec security.Securizer, service service.DummyService) {
	app.Get(root+"/:id", sec.Secure(keycloakRoles), handlers.GetDummyById(service))
	app.Get(root+"/external/api", sec.Secure(keycloakRoles), handlers.CallExternalApi(service))
	app.Get(root, sec.Secure(keycloakRoles), handlers.GetDummies(service))
	app.Post(root, sec.Secure(keycloakRoles), handlers.CreateDummy(service))
}
