package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
