package pedidorouters

import (
	"encoding/json"
	"net/http"
	"time"

	pedidobd "github.com/ascendere/resources/bd/pedido_bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
)

func DevolverPedido(w http.ResponseWriter, r *http.Request) {

	var t pedidomodels.Pedido

	err := json.NewDecoder(r.Body).Decode(&t)

	t.FechaDevolucion = time.Now()

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = pedidobd.DevuelvoPedido(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar devolver el pedido"+err.Error(), 400)
		return
	}

	if !status && err == nil {
		http.Error(w, "No ingreso correctamente el informe de devolución", http.StatusBadRequest)
	}


	if !status {
		http.Error(w, "No se ha logrado devolver el pedido", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}