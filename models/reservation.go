package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Phone    string             `bson:"phone"`
	Date     string             `bson:"date"`
	Time     string             `bson:"time"`
	Duration string             `bson:"duration"`
}
