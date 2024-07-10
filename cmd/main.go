package main

import (
	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/javiorfo/go-microservice/adapter/in/routes"
	"github.com/javiorfo/go-microservice/domain/service/dummy"
)

var (
	client       *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
)

func init() {
	client = gocloak.NewClient("http://localhost:8081")
	realm = "orfosys"
	clientID = "java-spring3-microservice"
	clientSecret = "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"
}

func SecureEndpoint(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
	}

	token := authHeader[len("Bearer "):]
	rptResult, err := client.RetrospectToken(c.Context(), token, clientID, clientSecret, realm)
	if err != nil || !*rptResult.Active {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	return c.JSON(fiber.Map{"message": "You have accessed a protected endpoint!"})
}

func main() {
    dummyService := dummy.DummyService{}
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
    api := app.Group("/orfosys")
    routes.DummyRouter(api, &dummyService)

// 	app.Get("/orfosys/dummy", SecureEndpoint)

	log.Fatal(app.Listen(":8080"))
}

