package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	RolId string `bson:"rolid" json:"rolid,omitempty"`
	Nombre string `bson:"nombre" json:"nombre,omitempty"`
	Apellidos string `bson:"apellidos" json:"apellidos,omitempty"`
	jwt.StandardClaims
}