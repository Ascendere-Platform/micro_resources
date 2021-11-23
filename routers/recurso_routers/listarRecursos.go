package recursorouters

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/resources/bd/recurso_bd"
)

func ListarRecursos(w http.ResponseWriter, r *http.Request) {

	result, status := recursobd.ListoRecursos()
	if !status {
		http.Error(w, "Error al leer los recursos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}