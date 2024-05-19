package repository_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/repository"
	database "github.com/unawaretub86/MoneyTracker/internal/domain/users/repository/database/mocks"
	"github.com/unawaretub86/MoneyTracker/utils/generic"
)

func Test_AllMethods(t *testing.T) {
	ass := assert.New(t)

	db := database.NewMockDatabase()

	repo := repository.NewMockUserRepository(db)

	testError := errors.New("test error")

	var id uint = 1
	name := "jhon doe"
	email := "alice.smith@example.com"
	birthDate := "1990-01-15"
	password := "password123"

	user := &entities.User{
		ID:        1,
		Name:      &name,
		Email:     &email,
		BirthDate: &birthDate,
		Password:  &password,
	}

	tests := []struct {
		funcName  string
		params    []interface{}
		patchData []interface{}
	}{
		{
			funcName: "GetUsers",
			params:   []interface{}{},
			patchData: []interface{}{
				nil,
				testError,
			},
		},
		{
			funcName: "GetUserByID",
			params: []interface{}{
				id,
			},
			patchData: []interface{}{
				user,
				testError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("method '%s' retrieves ok", tt.funcName), func(t *testing.T) {
			_, err := generic.CallMethod(db, fmt.Sprintf("Patch%s", tt.funcName), tt.patchData...)
			if err != nil {
				panic(err.Error())
			}

			result, err := generic.CallMethod(repo, tt.funcName, tt.params...)
			if err != nil {
				panic(err.Error())
			}

			errResult := result[1].(error)

			ass.NotNil(errResult)
			ass.Equal(errResult.Error(), testError.Error())
		})
	}
}
