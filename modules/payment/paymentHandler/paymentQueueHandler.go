package paymentHandler

import (
	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct {
		paymentUsecase paymentUsecase.PaymentUsecaseService
		cfg            *config.Config
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
