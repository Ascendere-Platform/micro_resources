package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoPedidos() ([]*pedidomodels.Pedido, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("pedido")

	var results []*pedidomodels.Pedido

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var pedido pedidomodels.Pedido
		err := cur.Decode(&pedido)
		if err != nil {
			return results, false
		}
		registro := make(map[string]interface{})
		if time.Now().After(pedido.TiempoPedido) && pedido.FechaDevolucion.IsZero(){
			pedido.Mensaje = "NO ENTREGADO, ATRASADO PARA DEVOLUCIÓN"
	
			registro["mensaje"] = pedido.Mensaje
			updtString := bson.M{
				"$set": registro,
			}
		
			filtro := bson.M{"_id": bson.M{"$eq": pedido.ID}}
		
			_, err := col.UpdateOne(ctx, filtro, updtString)
	
			if err != nil {
				return results, false
			}
		}
		results = append(results, &pedido)
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true

}
