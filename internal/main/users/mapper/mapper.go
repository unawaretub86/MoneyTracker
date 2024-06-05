package mapper

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/unawaretub86/MoneyTracker/internal/main/users/entities"
	"github.com/unawaretub86/MoneyTracker/internal/main/users/http/dto"
)

func CreateUserRqToUserEntity(req dto.CreateUserRequest) entities.User {
	now := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("Error parsing password")
		// Handle error
	}
	hashedPassString := string(hashedPassword)
	return entities.User{
		Name:      &req.Name,
		Username:  &req.Username,
		Email:     &req.Email,
		Password:  &hashedPassString,
		BirthDate: &req.BirthDate,
		CreatedAt: now,
	}
}
