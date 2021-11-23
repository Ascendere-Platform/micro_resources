package pedidorouters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pedidobd "github.com/ascendere/resources/bd/pedido_bd"
	pedidomodels "github.com/ascendere/resources/models/pedido_models"
	"github.com/ascendere/resources/routers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroPedido(w http.ResponseWriter, r *http.Request) {

	var t pedidomodels.Pedido
	err := json.NewDecoder(r.Body).Decode(&t)

	objID,_ := primitive.ObjectIDFromHex(routers.IDUsuario)

	t.Usuario.UsuarioID = objID
	t.Usuario.Email = routers.Email
	t.Usuario.Nombre = routers.Nombre
	t.TiempoPedido = t.FechaPedido.Add(time.Hour*120)
	t.Mensaje = "A TIEMPO PARA ENTREGAR"

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.InformePedido) == 0 {
		http.Error(w, "Informe de Pedido requerido ", 400)
		return
	}

	for _, recurso := range t.Recurso { 
		recursoEncontrado, err, mensaje := pedidobd.ChequeoExistenRecursos(recurso.RecursoID, recurso.CantidadPedida)

		if err != nil {
			http.Error(w, mensaje + " " + err.Error(), http.StatusBadRequest)
		}
		recurso.NombreRecurso = recursoEncontrado.NombreRecurso
	}

	t.FechaPedido = time.Now()

	status, err := pedidobd.RegistroPedido(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el pedido: "+ status + " "+err.Error(), 400)
		return
	}

	fmt.Sprintln(t)

	w.WriteHeader(http.StatusCreated)
}