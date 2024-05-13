package itemUsecase

import "github.com/liangkhwai/go-shop/modules/item/itemRepository"

type (
	ItemUsecaseService interface {
	}

	itemUsecasae struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecasae{
		itemRepository: itemRepository,
	}
}