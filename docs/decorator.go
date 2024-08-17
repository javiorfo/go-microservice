package docs

import (
	"strings"

	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice-lib/env"
	"github.com/swaggo/swag"
)

type SwaggerInfoWrapper struct {
	swag.Spec
}

const keycloakEnvVar = "KEYCLOAK_HOST"
const keycloakLocal = "http://localhost:8081"

func (i *SwaggerInfoWrapper) ReadDoc() string {
	i.BasePath = config.AppContextPath
	i.Title = config.AppName
	i.SwaggerTemplate = strings.Replace(docTemplate, keycloakEnvVar, env.GetEnvOrDefault(keycloakEnvVar, keycloakLocal), 1)
	return i.Spec.ReadDoc()
}
