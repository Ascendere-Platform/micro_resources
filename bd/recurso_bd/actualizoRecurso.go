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
	
	var resultado recursomodels.Recurso

	error := col.FindOne(ctx, bson.M{"_id":u.ID}).Decode(&resultado)

	if error != nil {
		return false, error
	}

	registro := make(map[string]interface{})

	if len(u.NombreRecurso) > 0{
		registro["nombreRecurso"] = u.NombreRecurso
	}

	if u.CantidadDisponible > 0 {
		registro["cantidadDisponible"] = u.CantidadDisponible
	}
	
	if resultado.CantidadExistente + u.CantidadExistente < 0 {
		return false, nil
	}

	if u.CantidadExistente != 0 {
		registro["cantidadExistente"] = resultado.CantidadExistente + u.CantidadExistente
		
		if resultado.CantidadDisponible + u.CantidadExistente < 0 {
			registro["cantidadDisponible"] = registro["cantidadExistente"]
		} else {
			registro["cantidadDisponible"] = resultado.CantidadDisponible + u.CantidadExistente
		}
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