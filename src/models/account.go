package models

type Account struct {
	Name      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt uint   `json:"created_at"`
}
