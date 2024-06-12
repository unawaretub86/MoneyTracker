package usecase

import (
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

func (useUser useCase) GetUserByID(id uint) (*entities.User, error) {
	return useUser.repo.GetUserByID(id)
}

func (useUser useCase) GetUsers() (*entities.Users, error) {
	return useUser.repo.GetUsers()
}

func (useUser useCase) UpdateUser(id uint, user *dto.UpdateUserRequest) (*entities.User,error)  {
	return useUser.repo.UpdateUser(id, user)
}
