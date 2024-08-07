package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/adapter/api/handlers"
	"github.com/javiorfo/go-microservice/adapter/persistence/repository"
	"github.com/javiorfo/go-microservice/common/database"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/service"
)

func Inject(api fiber.Router) {
    db := database.DBinstance
    dummyRepository := repository.NewDummyRepository(db)
    dummyService := service.NewDummyService(dummyRepository)

	handlers.DummyHandler(api, config.KeycloakConfig, dummyService)
}
