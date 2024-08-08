package config

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/javiorfo/go-microservice/internal/database"
	"github.com/javiorfo/go-microservice/internal/security"
)

// Keycloak configuration
var KeycloakConfig = security.KeycloakConfig{
	Keycloak:     gocloak.NewClient(getEnvOrDefault("KEYCLOAK_HOST", "http://localhost:8081")),
	Realm:        "javi",
	ClientID:     "srv-client",
	ClientSecret: getEnvOrDefault("KEYCLOAK_CLIENT_SECRET", "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"),
	Enabled:      false,
}

var DBDataConnection = database.DBDataConnection{
	Host:        getEnvOrDefault("DB_HOST", "localhost"),
	Port:        getEnvOrDefault("DB_PORT", "5432"),
	DBName:      getEnvOrDefault("DB_NAME", "db_dummy"),
	User:        getEnvOrDefault("DB_USER", "admin"),
	Password:    getEnvOrDefault("DB_PASSWORD", "admin"),
	ShowSQLInfo: true,
}

// App configuration
const AppName = "go-microservice"
const AppPort = ":8080"
const AppContextPath = "/app"

// Tracing server configuration
var TracingHost = getEnvOrDefault("TRACING_HOST", "localhost:4318")

// Swagger configuration
var SwaggerEnabled = getEnvOrDefault("SWAGGER_ENABLED", "true")

func getEnvOrDefault(envVar, fallback string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		return fallback
	}
	return value
}
