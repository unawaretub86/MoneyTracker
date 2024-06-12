package dto

type UpdateUserRequest struct {
	Name      *string   `json:"name"`
	BirthDate *string   `json:"birth_date"`
	Password  *string   `json:"password"`
}
