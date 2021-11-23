package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ActualizoRecurso(u recursomodels.Recurso) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	registro := make(map[string]interface{})

	if len(u.NombreRecurso) > 0{
		registro["nombreRecurso"] = u.NombreRecurso
	}
	if u.CantidadExistente > 0 {
		registro["cantidadExistente"] = u.CantidadExistente
	}
	if len(u.Imagen) > 0 {
		registro["imagen"] = u.Imagen
	}
	if len(u.TipoID) > 0 {
		registro["tipoid"] = u.TipoID
	}

	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}