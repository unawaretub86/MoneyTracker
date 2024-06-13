package dto

type CreateUserRequest struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	Password  string `json:"password"`
}
