package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func ListarTiposRecurso(w http.ResponseWriter, r *http.Request) {

	result, status := recursobd.ListoTipoRecurso()
	if !status {
		http.Error(w, "Error al leer los tipos de recurso", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}