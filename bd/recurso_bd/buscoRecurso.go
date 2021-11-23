package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoRecurso(id string) (recursomodels.DevuelvoRecurso, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")
	colTipo := db.Collection("tipoRecurso")

	var resultadoRecurso recursomodels.Recurso
	var resultadoTipo recursomodels.TipoRecurso

	var result recursomodels.DevuelvoRecurso

	objID, _ := primitive.ObjectIDFromHex(id)

	errRecurso := col.FindOne(ctx, bson.M{"_id":objID}).Decode(&resultadoRecurso)
	if errRecurso != nil {
		return result, errRecurso
	}

	errTipo := colTipo.FindOne(ctx, bson.M{"_id":resultadoRecurso.TipoID}).Decode(&resultadoTipo)

	if errTipo != nil {
		return result, errTipo
	}

	result.ID = resultadoRecurso.ID
	result.NombreRecurso = resultadoRecurso.NombreRecurso
	result.Imagen = resultadoRecurso.Imagen
	result.CantidadDisponible = resultadoRecurso.CantidadDisponible
	result.CantidadExistente = resultadoRecurso.CantidadExistente
	result.TipoRecurso.ID = resultadoTipo.ID
	result.TipoRecurso.NombreTipo = resultadoTipo.NombreTipo
	result.TipoRecurso.DescripcionTipo = resultadoTipo.DescripcionTipo

	return result, nil
}