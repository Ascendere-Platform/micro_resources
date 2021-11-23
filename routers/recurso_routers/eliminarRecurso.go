package recursorouters

import (
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func EliminarRecurso(w http.ResponseWriter, r *http.Request) {

	recursoID := r.URL.Query().Get("id")

	if len(recursoID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := recursobd.EliminoRecurso(recursoID)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}