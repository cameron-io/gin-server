package models

type Account struct {
	Name      string `json:"username" validate:"required,min=5,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=10,max=50"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
