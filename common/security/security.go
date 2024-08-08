package security

import (
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
	"github.com/javiorfo/go-microservice/common/response"
	"github.com/javiorfo/go-microservice/common/tracing"
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

var authorizationHeaderError = response.NewRestResponseError(response.ResponseError{Code: "AUTH", Message: "Authorization header missing"}) 
var invalidTokenError = response.NewRestResponseError(response.ResponseError{Code: "AUTH", Message: "Invalid or expired token"})
var unauthorizedRoleError = response.NewRestResponseError(response.ResponseError{Code: "AUTH", Message: "User does not have permission to access"})

func (kc KeycloakConfig) SecureWithRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Infof("%s Keycloak capture: %s", tracing.LogTraceAndSpan(c), c.Path())
		if !kc.Enabled {
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Errorf("%s %s", tracing.LogTraceAndSpan(c), authorizationHeaderError)
			return c.Status(http.StatusUnauthorized).JSON(authorizationHeaderError)
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		rptResult, err := kc.Keycloak.RetrospectToken(c.Context(), token, kc.ClientID, kc.ClientSecret, kc.Realm)
		if err != nil || !*rptResult.Active {
			log.Errorf("%s %s", tracing.LogTraceAndSpan(c), invalidTokenError)
			return c.Status(http.StatusUnauthorized).JSON(invalidTokenError)
		}

		if !userHasRole(kc.ClientID, token, roles) {
			log.Errorf("%s %s", tracing.LogTraceAndSpan(c), unauthorizedRoleError)
			return c.Status(http.StatusUnauthorized).JSON(unauthorizedRoleError)
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
