package usecase

import (
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/repository"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type (
	UseUser interface {
		GetUsers() (*entities.Users, error)
		GetUserByID(uint) (*entities.User, error)
	}

	useCase struct {
		repo repository.Repo
	}
)

func NewUseUser(container *dependencies.Container) UseUser {
	return &useCase{
		repo: repository.NewUserRepository(container),
	}
}
