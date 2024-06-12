package mock

import (
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

const (
	GetUsersMethodName    = "GetUsers"
	GetUserByIDMethodName = "GetUserByID"
	UpdateUserMethodName  = "UpdateUser"
)

func (mock *mockContainer) PatchGetUsers(users entities.Users, err error) {
	mock.mocker.Patch(GetUsersMethodName, users, err)
}

func (mock *mockContainer) GetUsers() (*entities.Users, error) {
	return mock.getUsersAndErrorFromResult(GetUsersMethodName)
}
func (mock *mockContainer) PatchGetUserByID(user *entities.User, err error) {
	mock.mocker.Patch(GetUserByIDMethodName, user, err)
}

func (mock *mockContainer) GetUserByID(uint) (*entities.User, error) {
	return mock.getUser(GetUserByIDMethodName)
}

func (mock *mockContainer) PatchUpdateUser(user *dto.UpdateUserRequest, err error) {
	mock.mocker.Patch(UpdateUserMethodName, user, err)
}

func (mock *mockContainer) UpdateUser(uint,  *entities.User) (*entities.User, error) {
	return mock.getUser(UpdateUserMethodName)
}

func (mock *mockContainer) getUsersAndErrorFromResult(name string) (*entities.Users, error) {
	result := mock.mocker.Get(name)

	var users entities.Users = nil
	if result[0] != nil {
		users = result[0].(entities.Users)
	}

	var err error = nil
	if result[1] != nil {
		err = result[1].(error)
	}

	return &users, err
}

func (mock *mockContainer) getUser(name string) (*entities.User, error) {
	result := mock.mocker.Get(name)

	var user *entities.User = nil
	if result[0] != nil {
		user = result[0].(*entities.User)
	}

	var err error = nil
	if result[1] != nil {
		err = result[1].(error)
	}

	return user, err
}
