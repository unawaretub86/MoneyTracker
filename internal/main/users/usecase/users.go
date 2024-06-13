package usecase

import (
	"fmt"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/mapper"
)

func (useUser useCase) GetUserByID(id uint) (*entities.User, error) {
	return useUser.repo.GetUserByID(id)
}

func (useUser useCase) GetUsers() (*entities.Users, error) {
	return useUser.repo.GetUsers()
}

func (useUser useCase) CreateUser(userDto dto.CreateUserRequest) (bool, error) {
	fmt.Print("creating entity: ", userDto.Username)

	userEntity, err := mapper.CreateUserRqToUserEntity(userDto)
	if err != nil {
		return false, err
	}

	return useUser.repo.CreateUser(userEntity)
}
