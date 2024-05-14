package repository

import (
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/entities"
	database "github.com/unawaretub86/MoneyTracker/internal/domain/users/repository/database"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type (
	Repo interface {
		GetUsers() (*entities.Users, error)
		GetUserByID(uint) (*entities.User, error)
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
