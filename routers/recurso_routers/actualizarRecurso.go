package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
	recursomodels "github.com/ascendere/resources/models/recurso_models"
)

func ActualizarRecurso(w http.ResponseWriter, r *http.Request) {

	var t recursomodels.Recurso

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = recursobd.ActualizoRecurso(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el recurso"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el recurso", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}