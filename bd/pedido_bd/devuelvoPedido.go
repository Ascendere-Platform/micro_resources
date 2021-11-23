package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func DevuelvoPedido(u pedidomodels.Pedido) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("pedido")
	colRecurso := db.Collection("recurso")

	var resultadoRecurso recursomodels.Recurso

	registro := make(map[string]interface{})

	if len(u.InformeDevolucion) < 0 {
		return false, nil
	}

	registro["informeDevolucion"] = u.InformeDevolucion
	registro["fechaDevolucion"] = u.FechaDevolucion

	if u.FechaDevolucion.After(u.TiempoPedido) {
		registro["mensaje"] = "ENTREGADO ATRASADO"
	}

	registro["mensaje"] = "ENTREGADO A TIEMPO"

	for _,recurso := range u.Recurso {
		condicion := bson.M{"_id": recurso.RecursoID}

		error := colRecurso.FindOne(ctx, condicion).Decode(&resultadoRecurso)

		if error != nil {
			return false, error
		}

		resultadoRecurso.CantidadDisponible = resultadoRecurso.CantidadDisponible + recurso.CantidadPedida

		registroRecurso := make(map[string]interface{})

		registroRecurso["cantidadDisponible"] = resultadoRecurso.CantidadDisponible
		updtString := bson.M{
			"$set": registroRecurso,
		}
	
		filtro := bson.M{"_id": bson.M{"$eq": resultadoRecurso.ID}}
	
		_, errUpd := colRecurso.UpdateOne(ctx, filtro, updtString)

		if errUpd != nil {
			return false, errUpd
		}
	}

	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}