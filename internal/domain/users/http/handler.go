package http

import (
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/usecase"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

const suffixErr = "Error"

type (
	Handler struct {
		UseUser usecase.UseUser
	}
)

func NewUserHandler(container *dependencies.Container) *Handler {
	return &Handler{
		UseUser: usecase.NewUseUser(container),
	}
}
