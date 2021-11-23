package recursorouters

import (
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func EliminarTipoRecurso(w http.ResponseWriter, r *http.Request) {

	tipoID := r.URL.Query().Get("id")

	if len(tipoID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := recursobd.EliminoTipoRecurso(tipoID)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}