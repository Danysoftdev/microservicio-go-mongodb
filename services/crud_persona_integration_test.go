package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/danysoftdev/microservicio-go-mongodb/config"
	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/repositories"
	"github.com/danysoftdev/microservicio-go-mongodb/services"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCrearBuscarActualizarEliminarPersona(t *testing.T) {
	ctx := context.Background()

	// 1. Iniciar contenedor de MongoDB
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

	// 2. Obtener puerto dinámico y construir URI
	endpoint, err := mongoC.Endpoint(ctx, "")
	assert.NoError(t, err)

	t.Setenv("MONGO_URI", "mongodb://"+endpoint)
	t.Setenv("MONGO_DB", "testdb")
	t.Setenv("COLLECTION_NAME", "personas_test")

	// 3. Conectar a Mongo y cerrar después
	err = config.ConectarMongo()
	assert.NoError(t, err)

	repositories.SetCollection(config.Collection)

	// Asegurarse de que la colección esté vacía antes de cada prueba
	_, err = config.Collection.DeleteMany(context.Background(), bson.M{})
	assert.NoError(t, err)


	// 4. Inyectar el repositorio real al servicio
	services.SetPersonaRepository(repositories.RealPersonaRepository{})

	// 5. Crear persona de prueba
	persona := models.Persona{
		Documento: "12345",
		Nombre:    "Persona",
		Apellido:  "Prueba",
		Edad:      28,
		Correo:    "persona@prueba.com",
		Telefono:  "3001234567",
		Direccion: "Calle Falsa 123",
	}

	// Crear
	err = services.CrearPersona(persona)
	assert.NoError(t, err)

	// Buscar
	encontrada, err := services.BuscarPersonaPorDocumento(persona.Documento)
	assert.NoError(t, err)
	assert.Equal(t, "Persona", encontrada.Nombre)

	// Actualizar
	persona.Nombre = "Persona Actualizada"
	persona.Correo = "nuevo@correo.com"
	err = services.ModificarPersona(persona.Documento, persona)
	assert.NoError(t, err)

	actualizada, err := services.BuscarPersonaPorDocumento(persona.Documento)
	assert.NoError(t, err)
	assert.Equal(t, "Persona Actualizada", actualizada.Nombre)
	assert.Equal(t, "nuevo@correo.com", actualizada.Correo)

	// Eliminar
	err = services.BorrarPersona(persona.Documento)
	assert.NoError(t, err)

	// Confirmar eliminación
	_, err = services.BuscarPersonaPorDocumento(persona.Documento)
	assert.Error(t, err)
	assert.Equal(t, "persona no encontrada", err.Error())

	defer config.CerrarMongo()
}
