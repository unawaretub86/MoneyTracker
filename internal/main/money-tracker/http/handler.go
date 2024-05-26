package http

import (
	"github.com/unawaretub86/MoneyTracker/internal/main/money-tracker/usecase"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type (
	Handler struct {
		UseCase usecase.UseCase
	}
)

func NewHandler(container *dependencies.Container) *Handler {
	return &Handler{
		UseCase: usecase.NewUse(container),
	}
}
