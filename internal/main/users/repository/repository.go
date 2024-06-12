package repository

import (
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
	database "github.com/unawaretub86/MoneyTracker/internal/main/users/repository/database"
)

type (
	Repo interface {
		GetUsers() (*entities.Users, error)
		GetUserByID(uint) (*entities.User, error)
		UpdateUser(uint, *dto.UpdateUserRequest) (*entities.User, error)
	}

	repository struct {
		databaseUser database.Database
	}
)

func NewUserRepository(container *dependencies.Container) Repo {
	return &repository{
		databaseUser: database.NewUserDatabase(container),
	}
}

func NewMockUserRepository(database database.Database) Repo {
	return &repository{
		databaseUser: database,
	}
}
