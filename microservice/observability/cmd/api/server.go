package main

import (
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/client"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/factory"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/handler"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/route"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/service"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type application struct {
	Logger            *zap.Logger
	TracerProvider    trace.TracerProvider
	PrometheusFactory factory.PrometheusFactory
}

func initApplication(a *application) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Get("/metrics", adaptor.HTTPHandler(a.PrometheusFactory.InitHandler()))

	app.Use(
		otelfiber.Middleware(
			otelfiber.WithTracerProvider(
				a.TracerProvider,
			),
			otelfiber.WithNext(func(c *fiber.Ctx) bool {
				return c.Path() == "/metrics"
			}),
		))

	client := client.NewClient()

	dataService := service.NewDataService(
		a.Logger,
		client,
	)

	dataHandler := handler.NewDataHandler(
		a.Logger,
		dataService,
	)

	r := route.NewRoute(
		dataHandler,
	)

	r.SetupRoutes(&route.AppContext{
		App: app,
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "404 not found",
		})
	})

	return app
}
