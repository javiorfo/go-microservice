package security

import (
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

type Securizer interface {
	SecureWithRoles(roles ...string) fiber.Handler
}

type KeycloakConfig struct {
	Keycloak     *gocloak.GoCloak
	Realm        string
	ClientID     string
	ClientSecret string
	Enabled      bool
}

func (kc KeycloakConfig) SecureWithRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !kc.Enabled {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header missing"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		rptResult, err := kc.Keycloak.RetrospectToken(c.Context(), token, kc.ClientID, kc.ClientSecret, kc.Realm)
		if err != nil || !*rptResult.Active {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}

		if !userHasRole(kc.ClientID, token, roles) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "user does not have permission to access"})
		}
		return c.Next()
	}
}

type customClaims struct {
	ResourceAccess map[string]any `json:"resource_access"`
	jwt.StandardClaims
}

func userHasRole(clientID, tokenStr string, roles []string) bool {
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
		log.Errorf("Resource key %s not found", clientID)
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
