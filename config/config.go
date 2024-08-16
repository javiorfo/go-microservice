package config

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/javiorfo/go-microservice/internal/database"
	"github.com/javiorfo/go-microservice/internal/env"
	"github.com/javiorfo/go-microservice/internal/security"
)

// Keycloak configuration
var KeycloakConfig = security.KeycloakConfig{
	Keycloak:     gocloak.NewClient(env.GetEnvOrDefault("KEYCLOAK_HOST", "http://localhost:8081")),
	Realm:        "javi",
	ClientID:     "srv-client",
	ClientSecret: env.GetEnvOrDefault("KEYCLOAK_CLIENT_SECRET", "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"),
	Enabled:      true,
}

var DBDataConnection = database.DBDataConnection{
	Host:        env.GetEnvOrDefault("DB_HOST", "localhost"),
	Port:        env.GetEnvOrDefault("DB_PORT", "5432"),
	DBName:      env.GetEnvOrDefault("DB_NAME", "db_dummy"),
	User:        env.GetEnvOrDefault("DB_USER", "admin"),
	Password:    env.GetEnvOrDefault("DB_PASSWORD", "admin"),
	ShowSQLInfo: true,
}

// App configuration
const AppName = "go-microservice"
const AppPort = ":8080"
const AppContextPath = "/app"

// Tracing server configuration
var TracingHost = env.GetEnvOrDefault("TRACING_HOST", "http://localhost:4318")

// Swagger configuration
var SwaggerEnabled = env.GetEnvOrDefault("SWAGGER_ENABLED", "true")
