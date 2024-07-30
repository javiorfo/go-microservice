package main

import (
	"context"
	"fmt"
	"log"
	"os"
	// 	"time"

	"github.com/gofiber/swagger"

	// 	"go.opentelemetry.io/otel"
	// 	"go.opentelemetry.io/otel/attribute"
	/* 	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	   	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	   	"go.opentelemetry.io/otel/sdk/resource"
	   	"go.opentelemetry.io/otel/sdk/trace"
	   	semconv "go.opentelemetry.io/otel/semconv/v1.26.0" */

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/javiorfo/go-microservice/adapter/api/routes"
	dummyRepository "github.com/javiorfo/go-microservice/adapter/persistence/dummy"
	"github.com/javiorfo/go-microservice/common/tracing"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/service/dummy"
)

/* func startTracing() (*trace.TracerProvider, error) {
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
				semconv.ServiceNameKey.String(config.AppName),
			),
		),
	)

	otel.SetTracerProvider(tracerprovider)

	return tracerprovider, nil
} */

func main() {
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

	app := fiber.New(fiber.Config{
		AppName: config.AppName,
	})

	file, err := os.OpenFile(fmt.Sprintf("./%s.log", config.AppName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{Output: file}))
	app.Use(cors.New())

	api := app.Group(config.AppContextPath)

	// Service injection
	/* 	dummyService := dummy.NewService()
	   	routes.DummyRouter(api, config.KeycloakConfig, dummyService) */
	injection(api)

	app.Get(fmt.Sprintf("%s/swagger/*", config.AppContextPath), swagger.HandlerDefault) // default

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

	log.Fatal(app.Listen(":" + config.AppPort))
}

func injection(api fiber.Router) {
	routes.DummyRouter(api, config.KeycloakConfig, dummy.NewService(dummyRepository.NewRepository()))
}
