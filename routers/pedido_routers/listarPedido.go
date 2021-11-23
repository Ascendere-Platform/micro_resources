package pedidorouters

import (
	"encoding/json"
	"net/http"

	pedidobd "github.com/ascendere/resources/bd/pedido_bd"
)

func ListarPedidos(w http.ResponseWriter, r *http.Request) {

	result, status := pedidobd.ListoPedidos()
	if !status {
		http.Error(w, "Error al leer los tipos de recurso", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}