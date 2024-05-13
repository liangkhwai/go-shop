package authHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface {
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		cfg:         cfg,
		authUsecase: authUsecase,
	}
}
