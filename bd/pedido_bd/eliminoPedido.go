package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminoPedido(pedidoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("pedido")

	objID, _ := primitive.ObjectIDFromHex(pedidoID)

	condicion := bson.M{
		"_id":objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}