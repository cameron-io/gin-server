package models

type User struct {
	Name      string `json:"name"     validate:"min=3,max=50"`
	Email     string `json:"email"    validate:"required,email"`
	Password  string `json:"password" validate:"required,min=10,max=50"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
