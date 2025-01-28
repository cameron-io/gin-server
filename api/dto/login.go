package dto

type Login struct {
	Email string `json:"email"    validate:"required,email"`
}
