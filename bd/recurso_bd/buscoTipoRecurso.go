package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func BuscoTipoRecurso(nombre string) (recursomodels.TipoRecurso, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("tipoRecurso")

	condicion := bson.M{"nombreTipo":nombre}
	
	var resultado recursomodels.TipoRecurso

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil{
		return resultado, err
	}
	return resultado, err
}