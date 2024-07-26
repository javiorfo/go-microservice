package security

import (
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

var (
	keycloak     *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
)

func init() {
	keycloak = gocloak.NewClient("http://localhost:8081")
	realm = "javi"
	clientID = "srv-client"
	clientSecret = "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"
}

func SecureEndpoint(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
	}

	token := authHeader[len("Bearer "):]
	rptResult, err := keycloak.RetrospectToken(c.Context(), token, clientID, clientSecret, realm)
	if err != nil || !*rptResult.Active {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	return c.JSON(fiber.Map{"message": "You have accessed a protected endpoint!"})
}

func SecureEndpointWithRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		rptResult, err := keycloak.RetrospectToken(c.Context(), token, clientID, clientSecret, realm)
		if err != nil || !*rptResult.Active {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		if !userHasRole(token, roles) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "user does not have permission to access"})
		}
		return c.Next()
	}
}

type customClaims struct {
/* 	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"` */
	ResourceAccess map[string]any `json:"resource_access"`
	jwt.StandardClaims
}

func userHasRole(tokenStr string, roles []string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &customClaims{})
	if err != nil {
		log.Errorf("Error parsing token: %v", err)
		return false
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		log.Errorf("Error asserting claims")
		return false
	}

	resourceData, ok := claims.ResourceAccess[clientID]
	if !ok {
		log.Fatalf("Resource key %s not found", clientID)
		return false
	}

	resourceMap := resourceData.(map[string]any)
	clientRoles := resourceMap["roles"].([]any)
	if len(clientRoles) > 0 {
		for _, cr := range clientRoles {
			for _, r := range roles {
				if r == cr.(string) {
					return true
				}
			}
		}
		return false
	} 

	log.Info("No roles found for resource key", clientID)
	return false
}
