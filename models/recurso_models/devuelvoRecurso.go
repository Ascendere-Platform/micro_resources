package recursomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevuelvoRecurso struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreRecurso      string             `bson:"nombreRecurso" json:"nombreRecurso,omitempty"`
	CantidadExistente  int                `bson:"cantidadExistente" json:"cantidadExistente,omitempty"`
	CantidadDisponible int                `bson:"cantidadDisponible" json:"cantidadDisponible,omitempty"`
	Imagen             string             `bson:"imagen" json:"imagen,omitempty"`
	TipoRecurso        struct {
		ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		NombreTipo      string             `bson:"nombreTipo" json:"nombreTipo,omitempty"`
		DescripcionTipo string             `bson:"descripcionTipo" json:"descripcionTipo,omitempty"`
	}
}
