package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func BuscarTipoRecurso(w http.ResponseWriter, r *http.Request) {

	tipo := r.URL.Query().Get("tipo")

	informacion, err := recursobd.BuscoRecurso(tipo)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un tipo de recurso ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}