package playerHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config,playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{cfg : cfg, playerUsecase : playerUsecase}
}
