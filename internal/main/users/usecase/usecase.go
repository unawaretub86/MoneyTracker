package usecase

import (
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/repository"
)

type (
	UseUser interface {
		GetUsers() (*entities.Users, error)
		GetUserByID(uint) (*entities.User, error)
		UpdateUser(uint, *dto.UpdateUserRequest) (*entities.User, error)
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

func NewMockUseUser(repositoryUser repository.Repo) UseUser {
	return &useCase{
		repo: repositoryUser,
	}
}
