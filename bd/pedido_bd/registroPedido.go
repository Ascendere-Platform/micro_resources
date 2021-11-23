package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroPedido(r pedidomodels.Pedido) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("pedido")

	registro := pedidomodels.Pedido{
		ID:            primitive.NewObjectID(),
		FechaPedido:   r.FechaPedido,
		IdProyecto:    r.IdProyecto,
		InformePedido: r.InformePedido,
		Estado:        true,
		Recurso:       r.Recurso,
		Usuario:       r.Usuario,
		TiempoPedido:  r.TiempoPedido,
		Mensaje:       r.Mensaje,
	}

	if len(r.Recurso) < 0 {
		r.Estado = false
		return "No ha solicitado recursos", nil
	}

	if len(r.InformePedido) < 0 {
		r.Estado = false
		return "Debe ingresar un informe para realizar el pedido", nil
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), nil
}
