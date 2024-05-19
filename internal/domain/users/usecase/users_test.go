package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/MoneyTracker/internal/domain/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/repository"
	database "github.com/unawaretub86/MoneyTracker/internal/domain/users/repository/database/mocks"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/usecase"
	"github.com/unawaretub86/MoneyTracker/utils/generic"
)

func Test_Methods__Ok(t *testing.T) {
	ass := assert.New(t)

	db := database.NewMockDatabase()

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

	users := entities.Users{
		*user,
	}

	tests := []struct {
		funcName    string
		dbFuncNames []string
		params      []interface{}
		patchData   [][]interface{}
		asserts     func(result []interface{})
	}{
		{
			funcName:    "GetUsers",
			dbFuncNames: []string{"GetUsers"},
			params:      []interface{}{},
			patchData: [][]interface{}{
				{
					users,
					nil,
				},
			},
			asserts: func(result []interface{}) {
				users := result[0].(*entities.Users)
				err := result[1]

				ass.Nil(err)
				ass.NotNil(users)
			},
		},
		{
			funcName:    "GetUserByID",
			dbFuncNames: []string{"GetUserByID"},
			params: []interface{}{
				id,
			},
			patchData: [][]interface{}{
				{
					user,
					nil,
				},
			},
			asserts: func(result []interface{}) {
				user := result[0].(*entities.User)
				err := result[1]

				ass.Nil(err)
				ass.NotNil(user)
				ass.Equal(user.Name, &name)
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("method '%s' retrieves ok", tt.funcName), func(t *testing.T) {
			for idx, dbFuncName := range tt.dbFuncNames {
				_, err := generic.CallMethod(db, fmt.Sprintf("Patch%s", dbFuncName), tt.patchData[idx]...)
				if err != nil {
					panic(err.Error())
				}
			}
			
			repo := repository.NewMockUserRepository(db)

			uc := usecase.NewMockUseUser(repo)

			result, err := generic.CallMethod(uc, tt.funcName, tt.params...)
			if err != nil {
				panic(err.Error())
			}

			tt.asserts(result)
		})
	}
}
