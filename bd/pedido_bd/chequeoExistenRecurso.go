package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursobd "github.com/ascendere/resources/bd/recurso_bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ChequeoExistenRecursos(recursoPedido pedidomodels.RecursoPedido) (string, error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	var resultado recursomodels.Recurso
	var nombre string

	objID, _ := primitive.ObjectIDFromHex(recursoPedido.RecursoID)

	error := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&resultado)

	nombre = resultado.NombreRecurso

	if error == nil {
		return nombre, error, "No se encuentra el recurso: " + recursoPedido.RecursoID
	}

	if resultado.CantidadDisponible == 0 {
		return "", error, "El recurso no se encuentra disponible"
	}

	calc := resultado.CantidadDisponible - recursoPedido.CantidadPedida

	resultado.CantidadDisponible = calc
	resultado.CantidadExistente = 0

	_,errorActu := recursobd.ActualizoRecurso(resultado)

	if errorActu != nil {
		return "no esta actualizando", errorActu, "Hubo un error al actualizar los recursos"
	}
	

	return nombre, error, recursoPedido.RecursoID
}
