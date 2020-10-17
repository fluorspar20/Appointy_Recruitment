package models

// Participant is a model for any participant object
type Participant struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	RSVP  bool   `json:"rsvp" bson:"rsvp"`
}
