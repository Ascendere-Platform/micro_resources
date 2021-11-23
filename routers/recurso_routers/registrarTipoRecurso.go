package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
)

func RegistrarTipoRecurso(w http.ResponseWriter, r *http.Request){
	var tipo recursomodels.TipoRecurso

	err := json.NewDecoder(r.Body).Decode(&tipo)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+ err.Error(),400)
		return
	}

	_, status, err := recursobd.RegistroTipoRecurso(tipo)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo Tipo", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo Tipo", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}