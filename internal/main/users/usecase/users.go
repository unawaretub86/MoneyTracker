package usecase

import "github.com/unawaretub86/MoneyTracker/internal/main/users/entities"

func (useUser useCase) GetUserByID(id uint) (*entities.User, error) {
	return useUser.repo.GetUserByID(id)
}

func (useUser useCase) GetUsers() (*entities.Users, error) {
	return useUser.repo.GetUsers()
}