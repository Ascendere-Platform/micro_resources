package recursobd

import (
	"context"
	"time"

	"github.com/ascendere/resources/bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroTipoRecurso(r recursomodels.TipoRecurso) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Recursos")
	col := db.Collection("tipoRecurso")

	registro := recursomodels.TipoRecurso{
		ID:                 primitive.NewObjectID(),
		NombreTipo:      r.NombreTipo,
		DescripcionTipo:  r.DescripcionTipo,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}