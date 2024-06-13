package mapper

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

func CreateUserRqToUserEntity(req dto.CreateUserRequest) (*entities.User, error) {
	now := time.Now()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {		
		return nil, err
	}

	hashedPassString := string(hashedPassword)
	
	return &entities.User{
		Name:      &req.Name,
		Username:  &req.Username,
		Email:     &req.Email,
		Password:  &hashedPassString,
		BirthDate: &req.BirthDate,
		CreatedAt: now,
	}, nil
}
