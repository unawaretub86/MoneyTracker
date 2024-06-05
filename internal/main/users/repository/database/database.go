package database

import (
	database "github.com/unawaretub86/MoneyTracker/internal/infrastructure/configuration/database"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
)

type (
	Database interface {
		GetUsers() (*entities.Users, error)
		GetUserByID(uint) (*entities.User, error)
		CreateUser(entities.User) (bool, error)
	}

	databaseUser struct {
		db database.Database
	}
)

func NewUserDatabase(container *dependencies.Container) Database {
	return &databaseUser{
		db: container.Database(),
	}
}
