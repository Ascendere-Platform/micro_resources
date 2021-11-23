package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminoTipoRecurso(tipoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("tipoRecurso")

	objID, _ := primitive.ObjectIDFromHex(tipoID)

	condicion := bson.M{
		"_id":objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}