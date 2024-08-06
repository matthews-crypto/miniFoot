package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Utilisateur struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nom      string             `bson:"nom"`
	Prenom   string             `bson:"prenom"`
	Phone    string             `bson:"phone"`
	CIN      string             `bson:"cin"`
	Role     string             `bson:"role"`
	Password string             `bson:"password"`
}
