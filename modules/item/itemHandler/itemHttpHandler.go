package itemHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface {
	}

	itemHttpUsecasae struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpUsecasae{
		cfg:         cfg,
		itemUsecase: itemUsecase,
	}
}
