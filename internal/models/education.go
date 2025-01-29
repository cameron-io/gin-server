package models

type Education struct {
	School       string `json:"school" validate:"required"`
	Degree       string `json:"degree" validate:"required"`
	FieldOfStudy string `json:"fieldofstudy" validate:"required"`
	From         int64  `json:"from" validate:"required"`
	To           int64  `json:"to"`
	Current      bool   `json:"current"`
	Description  string `json:"description"`
}
