package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoRecursos() ([]*recursomodels.DevuelvoRecurso, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	var results []*recursomodels.DevuelvoRecurso

	var resultadoRecurso []*recursomodels.Recurso

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var s recursomodels.Recurso
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}
			resultadoRecurso = append(resultadoRecurso, &s)
	}

	for _, recurso := range resultadoRecurso {
		colTipo := db.Collection("tipoRecurso")
		var tipo recursomodels.TipoRecurso

		err := colTipo.FindOne(ctx,bson.M{"_id":recurso.ID}).Decode(&tipo)

		if err !=nil {
			return results, false
		}

		aux := recursomodels.DevuelvoRecurso{
			ID: recurso.ID,
			NombreRecurso: recurso.NombreRecurso,
			CantidadExistente: recurso.CantidadExistente,
			CantidadDisponible: recurso.CantidadDisponible,
			TipoRecurso: tipo,
		}

		results = append(results, &aux)
	}

	if err != nil{
		return results, false
	}
	return results, true
}