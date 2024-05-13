package inventoryHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
