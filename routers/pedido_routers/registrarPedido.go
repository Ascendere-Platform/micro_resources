package pedidorouters

import (
	"encoding/json"
	"log"
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
	t.FechaPedido = time.Now()
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

	var nombres []string

	for _, recurso := range t.Recurso {
		if recurso.CantidadPedida < 0 {
			http.Error(w, "La cantidad pedida no puede ser negativa", 400)
			return
		}
		nombreRecurso, err, mensaje := pedidobd.ChequeoExistenRecursos(recurso)

		if err != nil {
			http.Error(w, mensaje + " " + err.Error() + "" + recurso.RecursoID, http.StatusBadRequest)
			return
		}
		nombres = append(nombres, nombreRecurso)

		log.Println(recurso.NombreRecurso)
	}

	for i := 0; i < len(t.Recurso); i++ {
		t.Recurso[i].NombreRecurso = nombres[i]
	}

	log.Println(t)
	

	status, err := pedidobd.RegistroPedido(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el pedido: "+ status + " "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}