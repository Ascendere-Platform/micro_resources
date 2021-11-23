package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroRecurso(r recursomodels.Recurso) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("recurso")

	registro := recursomodels.Recurso{
		ID:  primitive.NewObjectID(),
		NombreRecurso: r.NombreRecurso,
		CantidadExistente: r.CantidadExistente,
		CantidadDisponible: r.CantidadExistente,
		Imagen: r.Imagen,
		TipoID: r.TipoID,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}