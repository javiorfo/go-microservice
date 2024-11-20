package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gofiber/swagger"
	"go.opentelemetry.io/otel"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/javiorfo/go-microservice/config"
	_ "github.com/javiorfo/go-microservice/docs"
	"github.com/javiorfo/go-microservice/internal/injection"
	"github.com/javiorfo/go-microservice-lib/tracing"
)

//	@contact.name							API Support
//	@contact.email							fiber@swagger.io
//	@license.name							Apache 2.0
//	@license.url							http://www.apache.org/licenses/LICENSE-2.0.html
//	@securityDefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								KEYCLOAK_HOST/realms/javi/protocol/openid-connect/token
//	@scopes.read							Grants read access
//	@scopes.write							Grants write access
func main() {
	// Database
	err := config.DBDataConnection.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	// Tracing
	traceProvider, err := tracing.StartTracing(config.TracingHost, config.AppName)
	if err != nil {
		log.Fatalf("traceprovider: %v", err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("traceprovider: %v", err)
		}
	}()

	_ = traceProvider.Tracer(config.AppName)

	// Fiber
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Use(cors.New())

	// Tracing in context
	app.Use(func(c *fiber.Ctx) error {
		tracer := otel.Tracer(config.AppName)
		ctx, span := tracer.Start(c.Context(), c.Path())
		defer span.End()

		c.SetUserContext(ctx)
		c.Locals("traceID", span.SpanContext().TraceID())
		c.Locals("spanID", span.SpanContext().SpanID())

		return c.Next()
	})
	log.Info("Tracing configured!")

	// Web logger
	file, err := os.OpenFile(fmt.Sprintf("./%s.log", config.AppName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} ${ip}${path} | response: ${status} ${error} | [traceID: ${locals:traceID}, spanID: ${locals:spanID}] - ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05.000000",
		Output:     iw,
	}))

	// Context Path
	api := app.Group(config.AppContextPath)

	injection.Inject(api)

	// Swagger
	if config.SwaggerEnabled {
		app.Get(fmt.Sprintf("%s/swagger/*", config.AppContextPath), swagger.New(swagger.Config{
			DeepLinking:  false,
			DocExpansion: "list",
			OAuth: &swagger.OAuthConfig{
				Realm:        config.KeycloakConfig.Realm,
				ClientId:     config.KeycloakConfig.ClientID,
				ClientSecret: config.KeycloakConfig.ClientSecret,
			},
		}))
	}

	log.Infof("Context path: %s", config.AppContextPath)
	log.Infof("Starting %s on port %s...", config.AppName, config.AppPort)
	log.Info("Server Up!")
	log.Fatal(app.Listen(config.AppPort))
}
