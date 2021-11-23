package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminoRecurso(recursoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	objID, _ := primitive.ObjectIDFromHex(recursoID)

	condicion := bson.M{
		"_id":objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}