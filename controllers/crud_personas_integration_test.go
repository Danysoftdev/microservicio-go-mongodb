package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/danysoftdev/microservicio-go-mongodb/config"
	"github.com/danysoftdev/microservicio-go-mongodb/controllers"
	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/repositories"
	"github.com/danysoftdev/microservicio-go-mongodb/services"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestEndpointsControllerIntegration(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:6.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp").WithStartupTimeout(20 * time.Second),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	defer mongoC.Terminate(ctx)

	endpoint, err := mongoC.Endpoint(ctx, "")
	assert.NoError(t, err)

	os.Setenv("MONGO_URI", "mongodb://"+endpoint)
	os.Setenv("MONGO_DB", "testdb")
	os.Setenv("COLLECTION_NAME", "personas_test")

	err = config.ConectarMongo()
	assert.NoError(t, err)
	defer config.CerrarMongo()

	repositories.SetCollection(config.Collection)
	services.SetPersonaRepository(repositories.RealPersonaRepository{})

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/personas", controllers.CrearPersona).Methods("POST")
	router.HandleFunc("/personas", controllers.ObtenerPersonas).Methods("GET")
	router.HandleFunc("/personas/{documento}", controllers.ObtenerPersonaPorDocumento).Methods("GET")
	router.HandleFunc("/personas/{documento}", controllers.ActualizarPersona).Methods("PUT")
	router.HandleFunc("/personas/{documento}", controllers.EliminarPersona).Methods("DELETE")

	// 1. Crear persona
	persona := models.Persona{
		Documento: "999",
		Nombre:    "Test",
		Apellido:  "Integration",
		Edad:      33,
		Correo:    "test@integration.com",
		Telefono:  "1111111111",
		Direccion: "Calle Test",
	}
	body, _ := json.Marshal(persona)
	reqCrear := httptest.NewRequest("POST", "/personas", bytes.NewReader(body))
	resCrear := httptest.NewRecorder()
	router.ServeHTTP(resCrear, reqCrear)

	assert.Equal(t, http.StatusCreated, resCrear.Code)

	// 2. Obtener todas
	reqObtener := httptest.NewRequest("GET", "/personas", nil)
	resObtener := httptest.NewRecorder()
	router.ServeHTTP(resObtener, reqObtener)

	assert.Equal(t, http.StatusOK, resObtener.Code)
	content, _ := io.ReadAll(resObtener.Body)
	assert.Contains(t, string(content), "Test")

	// 3. Obtener por documento
	reqBuscar := httptest.NewRequest("GET", "/personas/999", nil)
	reqBuscar = mux.SetURLVars(reqBuscar, map[string]string{"documento": "999"})
	resBuscar := httptest.NewRecorder()
	router.ServeHTTP(resBuscar, reqBuscar)

	assert.Equal(t, http.StatusOK, resBuscar.Code)

	// 4. Modificar persona
	persona.Nombre = "Actualizado"
	bodyUpdate, _ := json.Marshal(persona)
	reqUpdate := httptest.NewRequest("PUT", "/personas/999", bytes.NewReader(bodyUpdate))
	reqUpdate = mux.SetURLVars(reqUpdate, map[string]string{"documento": "999"})
	resUpdate := httptest.NewRecorder()
	router.ServeHTTP(resUpdate, reqUpdate)

	assert.Equal(t, http.StatusOK, resUpdate.Code)

	// 5. Eliminar
	reqDelete := httptest.NewRequest("DELETE", "/personas/999", nil)
	reqDelete = mux.SetURLVars(reqDelete, map[string]string{"documento": "999"})
	resDelete := httptest.NewRecorder()
	router.ServeHTTP(resDelete, reqDelete)

	assert.Equal(t, http.StatusOK, resDelete.Code)
}
