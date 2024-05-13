package repository

import (
	database "github.com/unawaretub86/MoneyTracker/internal/domain/money-tracker/repository/database"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

type (
	Repo interface {
	}

	repository struct {
		database database.Database
	}
)

func NewRepository(container *dependencies.Container) Repo {
	return &repository{
		database: database.NewDatabase(container),
	}
}
