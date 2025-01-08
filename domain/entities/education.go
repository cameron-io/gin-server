package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Education struct {
	School       string             `json:"school" validate:"required"`
	Degree       string             `json:"degree" validate:"required"`
	FieldOfStudy string             `json:"fieldofstudy" validate:"required"`
	From         primitive.DateTime `json:"from" validate:"required"`
	To           primitive.DateTime `json:"to"`
	Current      bool               `json:"current"`
	Description  string             `json:"description"`
}
