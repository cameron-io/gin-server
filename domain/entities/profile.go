package entities

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	User           uuid.UUID `json:"user"`
	Company        string    `json:"company"`
	Website        string    `json:"website"`
	Location       string    `json:"location"`
	Status         string    `json:"status" validate:"required"`
	Skills         []string  `json:"skills" validate:"required"`
	Bio            string    `json:"bio"`
	GitHubUsername string    `json:"githubusername"`
	Experience     []Experience
	Education      []Education
	Date           primitive.DateTime `json:"date"`
}
