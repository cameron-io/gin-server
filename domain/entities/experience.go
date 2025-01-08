package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	Title       string             `json:"title" validate:"required"`
	Company     string             `json:"company" validate:"required"`
	Location    string             `json:"location"`
	From        primitive.DateTime `json:"from" validate:"required"`
	To          primitive.DateTime `json:"to"`
	Current     bool               `json:"current"`
	Description string             `json:"description"`
}
