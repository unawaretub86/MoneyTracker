package repository

import "github.com/unawaretub86/MoneyTracker/internal/main/users/entities"

func (userRepo repository) GetUserByID(id uint) (*entities.User, error) {
	return userRepo.databaseUser.GetUserByID(id)
}

func (userRepo repository) GetUsers() (*entities.Users, error) {
	return userRepo.databaseUser.GetUsers()
}

func (userRepo repository) CreateUser(user entities.User) (bool, error) {
	return userRepo.databaseUser.CreateUser(user)
}
