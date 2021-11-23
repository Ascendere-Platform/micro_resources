package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/resources/middlew"
	pedidorouters "github.com/ascendere/resources/routers/pedido_routers"
	recursorouters "github.com/ascendere/resources/routers/recurso_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	//Llamada al CRUD de Recursos
	router.HandleFunc("/registroRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.RegistroRecurso))).Methods("POST")
	router.HandleFunc("/eliminarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.EliminarRecurso))).Methods("DELETE")
	router.HandleFunc("/actualizarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ActualizarRecurso))).Methods("PUT")
	router.HandleFunc("/buscarRecurso", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.BuscarRecurso))).Methods("GET")
	router.HandleFunc("/listarRecursos", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ListarRecursos))).Methods("GET")

	//Llamada al CRUD de Tipos de Recursos
	router.HandleFunc("/registrarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.RegistrarTipoRecurso))).Methods("POST")
	router.HandleFunc("/eliminarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.EliminarTipoRecurso))).Methods("DELETE")
	router.HandleFunc("/buscarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.BuscarTipoRecurso))).Methods("GET")
	router.HandleFunc("/listarTipo", middlew.ChequeoBD(middlew.ValidoJWT(recursorouters.ListarTiposRecurso))).Methods("GET")

	//Llamada al CRUD de Pedidos
	router.HandleFunc("/registrarPedido", middlew.ChequeoBD(middlew.ValidoJWT(pedidorouters.RegistroPedido))).Methods("POST")
	router.HandleFunc("/eliminarPedido", middlew.ChequeoBD(middlew.ValidoJWT(pedidorouters.EliminarPedido))).Methods("DELETE")
	router.HandleFunc("/buscarPedido", middlew.ChequeoBD(middlew.ValidoJWT(pedidorouters.BuscarPedido))).Methods("GET")
	router.HandleFunc("/listarPedido", middlew.ChequeoBD(middlew.ValidoJWT(pedidorouters.ListarPedidos))).Methods("GET")
	router.HandleFunc("/devolverPedido", middlew.ChequeoBD(middlew.ValidoJWT(pedidorouters.DevolverPedido))).Methods("PUT")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}