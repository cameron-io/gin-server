package entities

type Experience struct {
	Title       string `json:"title" validate:"required"`
	Company     string `json:"company" validate:"required"`
	Location    string `json:"location"`
	From        int64  `json:"from" validate:"required"`
	To          int64  `json:"to"`
	Current     bool   `json:"current"`
	Description string `json:"description"`
}
