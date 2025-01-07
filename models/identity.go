package models

type Identity struct {
	Name   string `json:"name"     validate:"min=3,max=50"`
	Email  string `json:"email"    validate:"required,email"`
	Avatar string `json:"avatar"`
}
