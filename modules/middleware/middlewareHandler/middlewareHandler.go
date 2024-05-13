package middlewareHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface {
	}

	middlewareHandler struct {
		cfg *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config,middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg: cfg,
		middlewareUsecase: middlewareUsecase,
	}
}