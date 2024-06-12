package database

import (
	"gorm.io/gorm"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

func (databaseUser *databaseUser) GetUserByID(id uint) (*entities.User, error) {
	user, result := databaseUser.getUser(id)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (databaseUser *databaseUser) GetUsers() (*entities.Users, error) {
	users := entities.Users{}

	result := databaseUser.db.Connection().Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (databaseUser *databaseUser) UpdateUser(id uint, user *dto.UpdateUserRequest) (*dto.UpdateUserRequest,error)  {
	userResult, resultData := databaseUser.getUser(id)
	if resultData.Error != nil {
		return nil, resultData.Error
	}

	result := databaseUser.db.Connection().
	Where("id = ?", userResult.ID).
	Updates(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (databaseUser *databaseUser) getUser(id uint) (*entities.User, *gorm.DB) {
	user := &entities.User{}

	result := databaseUser.db.Connection().First(&user, id)

	return user, result
}
