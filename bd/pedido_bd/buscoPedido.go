package pedidobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPedido(ID string) (pedidomodels.Pedido, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("pedido")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id":objID}
	
	var resultado pedidomodels.Pedido

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	registro := make(map[string]interface{})

	if time.Now().After(resultado.TiempoPedido) && resultado.FechaDevolucion.IsZero(){
		resultado.Mensaje = "NO ENTREGADO, ATRASADO PARA DEVOLUCIÓN"

		registro["mensaje"] = resultado.Mensaje
		updtString := bson.M{
			"$set": registro,
		}
	
		filtro := bson.M{"_id": bson.M{"$eq": resultado.ID}}
	
		_, err := col.UpdateOne(ctx, filtro, updtString)

		if err != nil {
			return resultado, err
		}
	}


	if err != nil{
		return resultado, err
	}

	return resultado, err
}