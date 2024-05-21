package models

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	CreatedAt uint   `json:"created_at"`
}
