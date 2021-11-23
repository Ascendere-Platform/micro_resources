package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func BuscarRecurso(w http.ResponseWriter, r *http.Request) {

	recurso := r.URL.Query().Get("recurso")

	informacion, err := recursobd.BuscoRecurso(recurso)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un recurso "+ err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}