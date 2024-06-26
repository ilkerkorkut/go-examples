package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilkerkorkut/go-examples/microservice/observability/config"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/factory"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/logging"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/metric"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	ctx := context.Background()

	cm := config.NewConfigurationManager()

	exp, err := factory.NewOTLPExporter(ctx, cm.GetOTLPConfig().OTLPEndpoint)
	if err != nil {
		log.Printf("failed to create OTLP exporter: %v", err)
		return
	}

	tp := factory.NewTraceProvider(exp, cm.GetAppConfig().Name)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	logger, err := logging.InitLogger(
		cm.GetAppConfig().LogLevel,
		cm.GetKafkaConfig().Brokers,
	)
	if err != nil {
		log.Printf("failed to initialize logger : %v", err)
		return
	}

	promFactory := factory.NewPrometheusFactory(
		metric.TestCounter,
	)

	app := initApplication(&application{
		Logger:            logger,
		TracerProvider:    tp,
		PrometheusFactory: promFactory,
	})

	go func() {
		if serveErr := app.Listen(fmt.Sprintf(":%s", cm.GetAppConfig().Port)); serveErr != nil {
			logger.Fatal("failed to start server")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
}
