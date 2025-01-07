package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Experience struct {
	Title       string             `json:"title" validate:"required"`
	Company     string             `json:"company" validate:"required"`
	Location    string             `json:"location"`
	From        primitive.DateTime `json:"from" validate:"required"`
	To          primitive.DateTime `json:"to"`
	Current     bool               `json:"current"`
	Description string             `json:"description"`
}

type Education struct {
	School       string             `json:"school" validate:"required"`
	Degree       string             `json:"degree" validate:"required"`
	FieldOfStudy string             `json:"fieldofstudy" validate:"required"`
	From         primitive.DateTime `json:"from" validate:"required"`
	To           primitive.DateTime `json:"to"`
	Current      bool               `json:"current"`
	Description  string             `json:"description"`
}

type Profile struct {
	User           primitive.ObjectID `json:"user"`
	Company        string             `json:"company"`
	Website        string             `json:"website"`
	Location       string             `json:"location"`
	Status         string             `json:"status" validate:"required"`
	Skills         []string           `json:"skills" validate:"required"`
	Bio            string             `json:"bio"`
	GitHubUsername string             `json:"githubusername"`
	Experience     []Experience
	Education      []Education
	Date           primitive.DateTime `json:"date"`
}
