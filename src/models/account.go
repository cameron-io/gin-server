package models

type Account struct {
	Name      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
