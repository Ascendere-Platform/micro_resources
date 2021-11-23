package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoRecursos() ([]*recursomodels.Recurso, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	var resultadoRecurso []*recursomodels.Recurso

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return resultadoRecurso, err
	}

	for cur.Next(ctx) {
		var s recursomodels.Recurso
		err := cur.Decode(&s)
		if err != nil {
			return resultadoRecurso, err
		}
		resultadoRecurso = append(resultadoRecurso, &s)
	}

	if err != nil {
		return resultadoRecurso, err
	}
	return resultadoRecurso, err
}
