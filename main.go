package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danysoftdev/microservicio-go-mongodb/config"
	"github.com/danysoftdev/microservicio-go-mongodb/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Conectamos a MongoDB
	config.ConectarMongo()

	// Creamos el enrutador
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/crear-personas", controllers.CrearPersona).Methods("POST")
	router.HandleFunc("/buscar-personas", controllers.ObtenerPersonas).Methods("GET")
	router.HandleFunc("/buscar-personas/{documento}", controllers.ObtenerPersonaPorDocumento).Methods("GET")
	router.HandleFunc("/actualizar-personas/{documento}", controllers.ActualizarPersona).Methods("PUT")
	router.HandleFunc("/eliminar-personas/{documento}", controllers.EliminarPersona).Methods("DELETE")

	// Puerto de escucha
	puerto := ":8080"
	fmt.Printf("🚀 Servidor escuchando en http://localhost%s\n", puerto)
	log.Fatal(http.ListenAndServe(puerto, router))
}
