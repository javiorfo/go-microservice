package main

import (
	"log"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	client       *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
	keycloakURL  string
)

func init() {
	client = gocloak.NewClient("http://localhost:8081")
	realm = "orfosys"
	clientID = "java-spring3-microservice"
	clientSecret = "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"
	keycloakURL = "http://localhost:8081"
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
	app := fiber.New()

	// Use logger middleware
	app.Use(logger.New())

	// Define routes
	app.Get("/orfosys/dummy", SecureEndpoint)

	log.Fatal(app.Listen(":8080"))
}

