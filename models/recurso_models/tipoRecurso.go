package recursomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type TipoRecurso struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreTipo      string             `bson:"nombreTipo" json:"nombreTipo,omitempty"`
	DescripcionTipo string             `bson:"descripcionTipo" json:"descripcionTipo,omitempty"`
}