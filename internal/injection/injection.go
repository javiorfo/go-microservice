package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice-lib/integration"
	"github.com/javiorfo/go-microservice/api/routes"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/repository"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/internal/database"
)

func Inject(api fiber.Router) {
	// Database
	db := database.DBinstance
	dbIntegration := integration.DBinstance

	// Dummy: Repository, Servicer and Routes
	dummyRepository := repository.NewDummyRepository(db)
	async := integration.NewAsyncHttpClient(dbIntegration.Collection("async_dummies"), 3)
	dummyService := service.NewDummyService(dummyRepository, integration.NewHttpClient[integration.RawData](), async)
	routes.Dummy(api, config.KeycloakConfig, dummyService)
}
