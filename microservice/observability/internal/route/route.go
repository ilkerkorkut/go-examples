package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/handler"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/middleware"
)

type AppContext struct {
	App *fiber.App
}

type Route interface {
	SetupRoutes(ac *AppContext)
}

type route struct {
	dataHandler handler.DataHandler
}

func NewRoute(
	dataHandler handler.DataHandler,
) Route {
	return &route{
		dataHandler: dataHandler,
	}
}

func (r *route) SetupRoutes(ac *AppContext) {
	api := ac.App.Group("/",
		middleware.MetricsMiddleware(),
	)

	r.dataRoutes(api)
}

func (r *route) dataRoutes(fr fiber.Router) {
	fr.Get("/data/:id", r.dataHandler.Get)
}
