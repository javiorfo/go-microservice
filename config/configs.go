package config

import (

	"github.com/Nerzal/gocloak/v13"
	"github.com/javiorfo/go-microservice/common/security"
)

var KeycloakConfig = security.KeycloakConfig{
	Keycloak:     gocloak.NewClient("http://localhost:8081"),
	Realm:        "javi",
	ClientID:     "srv-client",
	ClientSecret: "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo",
	Enabled:      true,
}

const AppName = "go-microservice"
const AppPort = "8080"
const AppContextPath = "/app"

const TracingHost = "localhost:4318"
