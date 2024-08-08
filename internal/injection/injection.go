package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/handlers"
	"github.com/javiorfo/go-microservice/internal/database"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/domain/repository"
)

func Inject(api fiber.Router) {
    db := database.DBinstance
    dummyRepository := repository.NewDummyRepository(db)
    dummyService := service.NewDummyService(dummyRepository)

	handlers.DummyHandler(api, config.KeycloakConfig, dummyService)
}
