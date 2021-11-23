package pedidomodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecursoPedido struct {
	RecursoID      primitive.ObjectID `bson:"recursoid,omitempty" json:"recursoid"`
	NombreRecurso  string       `bson:"nombreRecurso" json:"nombreRecurso,omitempty"`
	CantidadPedida int          `bson:"cantidadpedida" json:"cantidadpedida,omitempty"`
}