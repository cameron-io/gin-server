package entities

type User struct {
	Name      string `json:"name"     validate:"min=3,max=50"`
	Email     string `json:"email"    validate:"required,email"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
