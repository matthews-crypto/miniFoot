package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Terrain struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nom      string             `bson:"nom"`
	Prenom   string             `bson:"prenom"`
	Gerant   primitive.ObjectID `bson:"gerant"`
	Proprio  primitive.ObjectID `bson:"proprio"`
	Position string             `bson:"position"`
}
