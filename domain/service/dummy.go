package service

import (
	"context"
	"net/http"

	"github.com/javiorfo/go-microservice-lib/integration"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type DummyService interface {
	FindById(context.Context, uint) (*model.Dummy, error)
	FindAll(context.Context, pagination.Page, string) ([]model.Dummy, error)
	Create(context.Context, *model.Dummy) error
	External(context.Context) (string, error)
}

type dummyService struct {
	repository repository.DummyRepository
	client     integration.Client[integration.RawData]
	async      integration.Async
	tracer     trace.Tracer
}

func NewDummyService(r repository.DummyRepository, c integration.Client[integration.RawData], a integration.Async) DummyService {
	return &dummyService{
		repository: r,
		client:     c,
		async:      a,
		tracer:     otel.Tracer(tracing.Name()),
	}
}

func (service *dummyService) FindById(ctx context.Context, id uint) (*model.Dummy, error) {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.FindById(ctx, id)
}

func (service *dummyService) FindAll(ctx context.Context, page pagination.Page, info string) ([]model.Dummy, error) {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.FindAll(ctx, page, info)
}

func (service *dummyService) Create(ctx context.Context, d *model.Dummy) error {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.Create(ctx, d)
}

type data struct {
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (service *dummyService) External(ctx context.Context) (string, error) {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	service.async.Execute(integration.NewRequest(context.Background(), "https://jsonplaceholder.typicode.com/posts",
		integration.WithMethod(http.MethodPost),
		integration.WithBody(data{UserId: 100, Title: "test"}),
		integration.WithJsonHeaders(),
	))

	resp, err := service.client.Send(integration.NewRequest(ctx, "https://jsonplaceholder.typicode.com/todos/1", integration.WithJsonHeaders()))

	if err != nil {
		return "", err
	}

	v := resp.ValueFromJsonField("title").OrElse("no value")

	return v.(string), nil
}
