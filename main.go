package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danysoftdev/microservicio-go-mongodb/config"
	"github.com/danysoftdev/microservicio-go-mongodb/controllers"
	"github.com/danysoftdev/microservicio-go-mongodb/repositories"
	"github.com/danysoftdev/microservicio-go-mongodb/services"

	"github.com/gorilla/mux"
)

func main() {
	// Conectamos a MongoDB
	err := config.ConectarMongo()
	if err != nil {
		log.Fatal("❌ Error conectando a MongoDB:", err)
	}


	// 2. Inyectar el repositorio real
	services.SetPersonaRepository(repositories.RealPersonaRepository{})

	// 3. Inyectar la colección de MongoDB
	repositories.SetCollection(config.Collection)

	// Creamos el enrutador
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/crear-personas", controllers.CrearPersona).Methods("POST")
	router.HandleFunc("/listar-personas", controllers.ObtenerPersonas).Methods("GET")
	router.HandleFunc("/buscar-personas/{documento}", controllers.ObtenerPersonaPorDocumento).Methods("GET")
	router.HandleFunc("/actualizar-personas/{documento}", controllers.ActualizarPersona).Methods("PUT")
	router.HandleFunc("/eliminar-personas/{documento}", controllers.EliminarPersona).Methods("DELETE")

	// Puerto de escucha
	puerto := ":8080"
	fmt.Printf("🚀 Servidor escuchando en http://localhost%s\n", puerto)
	log.Fatal(http.ListenAndServe(puerto, router))
}
