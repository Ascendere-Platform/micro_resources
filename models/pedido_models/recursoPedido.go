package pedidomodels


type RecursoPedido struct {
	RecursoID      string `bson:"recursoid,omitempty" json:"recursoid"`
	NombreRecurso  string       `bson:"nombreRecurso" json:"nombreRecurso,omitempty"`
	CantidadPedida int          `bson:"cantidadpedida" json:"cantidadpedida,omitempty"`
}