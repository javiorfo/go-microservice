package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/swagger"

	"go.opentelemetry.io/otel"
	// 	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/javiorfo/go-microservice/adapter/in/routes"
	"github.com/javiorfo/go-microservice/domain/service/dummy"
)

var (
	client       *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
)

func init() {
	client = gocloak.NewClient("http://localhost:8081")
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
	rptResult, err := client.RetrospectToken(c.Context(), token, clientID, clientSecret, realm)
	if err != nil || !*rptResult.Active {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}
    return c.Next()
// 	return c.JSON(fiber.Map{"message": "You have accessed a protected endpoint!"})
}

func startTracing() (*trace.TracerProvider, error) {
	serviceName := "go-microservice"
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("error crearing exporter: %w", err)
	}

	tracerprovider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
// 				attribute.String("environment", "testing"),
			),
		),
	)

	otel.SetTracerProvider(tracerprovider)

	return tracerprovider, nil
}

func main() {
	// Tracing
    traceProvider, err := startTracing()
	if err != nil {
		log.Fatalf("traceprovider: %v", err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("traceprovider: %v", err)
		}
	}()

	_ = traceProvider.Tracer("my-app")

	dummyService := dummy.Service{}
	app := fiber.New(fiber.Config{
        AppName: "go-microservice v0.1.0",
    })

    file, err := os.OpenFile("./go-microservice.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer file.Close()
	app.Use(logger.New(logger.Config{ Output: file }))
	app.Use(cors.New())
//     app.Use("/app", SecureEndpoint)

	api := app.Group("/app")
	routes.DummyRouter(api, &dummyService)

	app.Get("/app/swagger/*", swagger.HandlerDefault) // default

    // /app/swagger/index.html
/* 	app.Get("/app/swagger/*", swagger.New(swagger.Config{ // custom
        URL:         "http://localhost:8080/app/doc.json",
		DeepLinking: false,
		DocExpansion: "none",
		OAuth: &swagger.OAuthConfig{
			AppName:  "srv-client",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	})) */

	// 	app.Get("/app/dummy", SecureEndpoint)

	log.Fatal(app.Listen(":8080"))
}
