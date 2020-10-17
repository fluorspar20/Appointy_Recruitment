package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Meeting is a model for any meeting object
type Meeting struct {
	// ID will be automatically generated using the mongo driver
	Title        string   `json:"title" bson:"title"`
	Participants []string `json:"participants" bson:"participants"` //the email-id of participants is stored in the array
	StartTime    time.Time
	EndTime      time.Time
	Timestamp    primitive.Timestamp
}
