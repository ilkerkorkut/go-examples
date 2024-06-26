package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/service"
	"go.uber.org/zap"
)

type DataHandler interface {
	Get(c *fiber.Ctx) error
}

type dataHandler struct {
	logger      *zap.Logger
	dataService service.DataService
}

func NewDataHandler(
	logger *zap.Logger,
	dataService service.DataService,
) DataHandler {
	return &dataHandler{
		logger:      logger,
		dataService: dataService,
	}
}

func (h *dataHandler) Get(c *fiber.Ctx) error {
	ctx := c.UserContext()

	id := c.Params("id")

	data, err := h.dataService.GetData(ctx, id)
	if err != nil {
		h.logger.Error("failed to get data from service", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get data from service",
		})
	}

	return c.JSON(data)
}
