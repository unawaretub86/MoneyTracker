package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/usecase"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
)

func Test_NewUseUser(t *testing.T) {
	ass := assert.New(t)

	uc := usecase.NewUseUser(&dependencies.Container{})

	ass.NotNil(uc)
}

func Test_NewUseMockUser(t *testing.T) {
	ass := assert.New(t)

	usecaseMock := usecase.NewUseUser(&dependencies.Container{})

	uc := usecase.NewMockUseUser(usecaseMock)

	ass.NotNil(uc)
}
