package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Director  string             `bson:"director,omitempty"`
	Rating    float32            `bson:"rating,omitempty"`
	CreatedOn time.Time          `bson:"createdon,omitempty"`
}
