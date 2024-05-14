package entities

import "time"

type (
	User struct {
		ID        int       `json:"id"`
		Name      *string   `json:"name"`
		Email     *string   `json:"email"`
		BirthDate *string   `json:"birth_date"`
		Password  *string   `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Users []User
)

func (User) TableName() string {
	return "mt_users"
}
