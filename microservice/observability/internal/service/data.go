package service

import (
	"context"
	"time"

	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/client"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type DataService interface {
	GetData(
		ctx context.Context,
		id string,
	) (map[string]any, error)
}

type dataService struct {
	logger *zap.Logger
	client client.Client
}

func NewDataService(
	logger *zap.Logger,
	client client.Client,
) DataService {
	return &dataService{
		logger: logger,
		client: client,
	}
}

func (s *dataService) GetData(
	ctx context.Context,
	id string,
) (map[string]any, error) {
	span := trace.SpanFromContext(ctx)

	span.AddEvent("Starting fake long running task")

	sleepTime := 70
	sleepTimeDuration := time.Duration(sleepTime * int(time.Millisecond))
	time.Sleep(sleepTimeDuration)

	span.AddEvent("Done first fake long running task")

	span.AddEvent("Starting request to Star Wars API")

	s.logger.
		With(zap.String("traceID", span.SpanContext().TraceID().String())).
		Info("Getting data", zap.String("id", id))
	res, err := s.client.StarWars().GetPerson(ctx, id)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent("Done request to Star Wars API")

	return res, nil
}
