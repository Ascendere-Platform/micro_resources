package pedidorouters

import (
	"net/http"

	pedidobd "github.com/ascendere/resources/bd/pedido_bd"
)

func EliminarPedido(w http.ResponseWriter, r *http.Request) {

	pedidoID := r.URL.Query().Get("id")

	if len(pedidoID) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := pedidobd.EliminoPedido(pedidoID)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}