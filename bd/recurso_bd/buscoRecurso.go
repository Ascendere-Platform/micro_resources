package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func BuscoRecurso(nombre string) (recursomodels.DevuelvoRecurso, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	condiciones := make([]bson.M,0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"nombreRecurso": nombre}})
	condiciones = append(condiciones, bson.M{
		"lookup": bson.M{
			"from":         "tipoRecurso",
			"localField":   "tipoid",
			"foreignField": "_id",
			"as":           "tipo",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$asignatura"})

	cursor, err := col.Aggregate(ctx, condiciones)

	var result recursomodels.DevuelvoRecurso

	err = cursor.All(ctx, &result)

	if err != nil{
		return result, err
	}
	return result, err
}