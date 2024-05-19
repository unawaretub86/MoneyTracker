package database_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/domain/users/repository/database"
	"github.com/unawaretub86/MoneyTracker/internal/infrastructure/dependencies"
	"github.com/unawaretub86/MoneyTracker/utils/generic"
)

func Test_AllMethods_Ok(t *testing.T) {
	ass := assert.New(t)

	var id uint = 1

	tests := []struct {
		funcName    string
		params      []interface{}
		mockQueries func(mock sqlmock.Sqlmock)
		asserts     func(result []interface{})
	}{
		{
			funcName: "GetUsers",
			params:   []interface{}{},
			mockQueries: func(mock sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows([]string{"id", "name", "email", "birthDate", "password"}).
					AddRow(1, "jhon doe", "alice.smith@example.com", "1990-01-15", "password123")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT *`)).
					WillReturnRows(userRows)
			},
			asserts: func(result []interface{}) {
				users := result[0].(*entities.Users)
				err := result[1]

				ass.Nil(err)
				ass.NotEmpty(users)
			},
		},
		{
			funcName: "GetUserByID",
			params: []interface{}{
				id,
			},
			mockQueries: func(mock sqlmock.Sqlmock) {
				userRows := sqlmock.NewRows([]string{"id", "name", "email", "birthDate", "password"}).
					AddRow(1, "jhon doe", "alice.smith@example.com", "1990-01-15", "password123")

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "mt_users" WHERE "mt_users"."id" = $1 ORDER BY "mt_users"."id" LIMIT 1`)).
					WithArgs(id).
					WillReturnRows(userRows)
			},
			asserts: func(result []interface{}) {
				user := result[0].(*entities.User)
				err := result[1]

				ass.Nil(err)
				ass.NotNil(user)
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("method '%s' retrieves ok", tt.funcName), func(*testing.T) {
			container, mock := dependencies.StartMockDependencies()

			databaseRepository := database.NewUserDatabase(container)

			tt.mockQueries(mock)

			result, err := generic.CallMethod(databaseRepository, tt.funcName, tt.params...)
			if err != nil {
				panic(err.Error())
			}

			tt.asserts(result)
		})
	}
}
