package repository

import (
	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

func (userRepo repository) GetUserByID(id uint) (*entities.User, error) {
	return userRepo.databaseUser.GetUserByID(id)
}

func (userRepo repository) GetUsers() (*entities.Users, error) {
	return userRepo.databaseUser.GetUsers()
}

func (userRepo repository) UpdateUser(id uint, req *dto.UpdateUserRequest) (*entities.User,error)  {
	

	return 
}
