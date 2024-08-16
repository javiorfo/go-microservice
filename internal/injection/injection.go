package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/routes"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/repository"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/internal/database"
)

func Inject(api fiber.Router) {
	// Database
	db := database.DBinstance

	// Dummy: Repository, Servicer and Routes
	dummyRepository := repository.NewDummyRepository(db)
	dummyService := service.NewDummyService(dummyRepository)
	routes.Dummy(api, config.KeycloakConfig, dummyService)
}
