package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChequeoExistenRecursos(id primitive.ObjectID, cantidadPedida int) (pedidomodels.RecursoPedido, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recurso")
	col := db.Collection("recurso")

	var resultado recursomodels.Recurso

	var resultadoRecursoPedido pedidomodels.RecursoPedido

	condicion := bson.M{"_id": id}

	error := col.FindOne(ctx, condicion).Decode(&resultado)
	if error != nil {
		return resultadoRecursoPedido, false, "No se encuentra el recurso"
	}

	if resultado.CantidadDisponible == 0 {
		return resultadoRecursoPedido, false, "El recurso no se encuentra disponible"
	}

	if cantidadPedida > resultado.CantidadDisponible {
		return resultadoRecursoPedido, false, "No se dispone de tantos recursos"
	}

	if cantidadPedida > resultado.CantidadExistente {
		return resultadoRecursoPedido, false, "No existen tantos recursos"
	}

	if cantidadPedida < resultado.CantidadDisponible && cantidadPedida > 0 {
		resultado.CantidadDisponible = resultado.CantidadDisponible - cantidadPedida

		registro := make(map[string]interface{})

		registro["cantidadDisponible"] = resultado.CantidadDisponible
		updtString := bson.M{
			"$set": registro,
		}
	
		filtro := bson.M{"_id": bson.M{"$eq": resultado.ID}}
	
		_, err := col.UpdateOne(ctx, filtro, updtString)

		if err != nil {
			return resultadoRecursoPedido, false, err.Error()
		}

		resultadoRecursoPedido.RecursoID = resultado.ID
		resultadoRecursoPedido.NombreRecurso = resultado.NombreRecurso
		resultadoRecursoPedido.CantidadPedida = cantidadPedida
	
	}

	return resultadoRecursoPedido, true, id.String()
}