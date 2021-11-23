package recursomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recurso struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreRecurso string       `bson:"nombreRecurso" json:"nombreRecurso,omitempty"`
	CantidadExistente        int       `bson:"cantidadExistente" json:"cantidadExistente,omitempty"`
	CantidadDisponible       int       `bson:"cantidadDisponible" json:"cantidadDisponible,omitempty"`
	Imagen          string       `bson:"imagen" json:"imagen,omitempty"`
	TipoID primitive.ObjectID `bson:"tipoid,omitempty" json:"tipoid"`
}