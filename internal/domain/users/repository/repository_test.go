package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/MoneyTracker/internal/domain/users/repository"
	database "github.com/unawaretub86/MoneyTracker/internal/domain/users/repository/database/mocks"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

func Test_NewRepositoryUsers(t *testing.T) {
	ass := assert.New(t)

	container := &dependencies.Container{}

	repo := repository.NewUserRepository(container)

	ass.NotNil(repo)
}

func Test_NewMockRepositoryUsers(t *testing.T) {
	ass := assert.New(t)

	databaseRepository := database.NewMockDatabase()

	repo := repository.NewMockUserRepository(databaseRepository)

	ass.NotNil(repo)
}
