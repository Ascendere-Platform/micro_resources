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

	condiciones := make([]bson.M,0)

	condiciones = append(condiciones, bson.M{})
	condiciones = append(condiciones, bson.M{
		"lookup": bson.M{
			"from":         "tipoRecurso",
			"localField":   "tipoid",
			"foreignField": "_id",
			"as":           "tipo",
		}})
	condiciones = append(condiciones,  bson.M{"$unwind": "$asignatura"})

	cursor, err := col.Aggregate(ctx, condiciones)

	err = cursor.All(ctx, &results)

	if err != nil{
		return results, false
	}
	return results, true
}