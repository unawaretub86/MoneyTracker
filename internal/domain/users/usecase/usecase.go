package usecase

import (
	"github.com/unawaretub86/MoneyTracker/internal/domain/money-tracker/repository"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type (
	UseCase interface {
	}

	useCase struct {
		repo repository.Repo
	}
)

func NewUse(container *dependencies.Container) UseCase {
	return &useCase{
		repo: repository.NewRepository(container),
	}
}
