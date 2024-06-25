package security

import (
	"net/http"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
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
