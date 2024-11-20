package config

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/javiorfo/go-microservice/internal/database"
	"github.com/javiorfo/go-microservice-lib/env"
	"github.com/javiorfo/go-microservice-lib/security"
)

// Keycloak configuration
var KeycloakConfig = security.KeycloakConfig{
	Keycloak:     gocloak.NewClient(env.GetEnvOr("KEYCLOAK_HOST", "http://localhost:8081")),
	Realm:        "javi",
	ClientID:     "srv-client",
	ClientSecret: env.GetEnvOr("KEYCLOAK_CLIENT_SECRET", "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"),
	Enabled:      true,
}

var DBDataConnection = database.DBDataConnection{
	Host:        env.GetEnvOr("DB_HOST", "localhost"),
	Port:        env.GetEnvOr("DB_PORT", "5432"),
	DBName:      env.GetEnvOr("DB_NAME", "db_dummy"),
	User:        env.GetEnvOr("DB_USER", "admin"),
	Password:    env.GetEnvOr("DB_PASSWORD", "admin"),
	ShowSQLInfo: true,
}

// App configuration
const AppName = "go-microservice"
const AppPort = ":8080"
const AppContextPath = "/app"

// Tracing server configuration
var TracingHost = env.GetEnvOr("TRACING_HOST", "http://localhost:4318")

// Swagger configuration
var SwaggerEnabled = env.GetEnvOr("SWAGGER_ENABLED", true)
