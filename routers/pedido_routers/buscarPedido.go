package pedidorouters

import (
	"encoding/json"
	"net/http"

	pedidobd "github.com/ascendere/resources/bd/pedido_bd"
)

func BuscarPedido(w http.ResponseWriter, r *http.Request) {

	pedido := r.URL.Query().Get("id")

	informacion, err := pedidobd.BuscoPedido(pedido)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un pedido ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}