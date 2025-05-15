package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice/domain/model"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type DummyRepository interface {
	FindById(context.Context, uint) (*model.Dummy, error)
	FindAll(context.Context, pagination.Page, string) ([]model.Dummy, error)
	Create(context.Context, *model.Dummy) error
}

type dummyRepository struct {
	*gorm.DB
	tracer trace.Tracer
}

func NewDummyRepository(db *gorm.DB) DummyRepository {
	return &dummyRepository{DB: db, tracer: otel.Tracer(tracing.Name())}
}

func (repository *dummyRepository) FindById(ctx context.Context, id uint) (*model.Dummy, error) {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	var dummy model.Dummy
	result := repository.WithContext(ctx).Find(&dummy, "id = ?", id)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Dummy not found")
	}

	return &dummy, nil
}

func (repository *dummyRepository) FindAll(ctx context.Context, page pagination.Page, info string) ([]model.Dummy, error) {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	var dummies []model.Dummy
	filter := repository.WithContext(ctx).
		Offset(page.Page - 1).
		Limit(page.Size).
		Order(fmt.Sprintf("%s %s", page.SortBy, page.SortOrder))

	if info != "" {
		filter = filter.Where("info = ?", info)
	}
	
	results := filter.Find(&dummies)

	if err := results.Error; err != nil {
		return nil, err
	}

	return dummies, nil
}

func (repository *dummyRepository) Create(ctx context.Context, d *model.Dummy) error {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	result := repository.DB.WithContext(ctx).Create(d)
	if err := result.Error; err != nil {
		return fmt.Errorf("Error creating dummy %v", err)
	}
	return nil
}
