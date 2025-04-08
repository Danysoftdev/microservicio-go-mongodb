package services_test

import (
	"errors"
	"testing"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockRepositorio implementa la interfaz PersonaRepository
type MockRepositorioBuscar struct {
	mock.Mock
}

func (m *MockRepositorioBuscar) InsertarPersona(p models.Persona) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockRepositorioBuscar) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}

func (m *MockRepositorioBuscar) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockRepositorioBuscar) ActualizarPersona(doc string, p models.Persona) error {
	args := m.Called(doc, p)
	return args.Error(0)
}

func (m *MockRepositorioBuscar) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}

func TestBuscarPersonaPorDocumento_Exito(t *testing.T) {
	mockRepo := new(MockRepositorioBuscar)
	services.SetPersonaRepository(mockRepo)

	personaMock := models.Persona{
		Documento: "123",
		Nombre:    "Juan",
		Apellido:  "Pérez",
		Edad:      30,
		Correo:    "juan@example.com",
		Telefono:  "1234567890",
		Direccion: "Calle 123",
	}

	mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(personaMock, nil)

	persona, err := services.BuscarPersonaPorDocumento("123")

	assert.Nil(t, err)
	assert.Equal(t, "Juan", persona.Nombre)
	mockRepo.AssertExpectations(t)
}

func TestBuscarPersonaPorDocumento_Vacio(t *testing.T) {
	mockRepo := new(MockRepositorioBuscar)
	services.SetPersonaRepository(mockRepo)

	_, err := services.BuscarPersonaPorDocumento("")

	assert.NotNil(t, err)
	assert.Equal(t, "el documento no puede estar vacío", err.Error())
}

func TestBuscarPersonaPorDocumento_NoEncontrado(t *testing.T) {
	mockRepo := new(MockRepositorioBuscar)
	services.SetPersonaRepository(mockRepo)

	mockRepo.On("ObtenerPersonaPorDocumento", "999").Return(models.Persona{}, mongo.ErrNoDocuments)

	_, err := services.BuscarPersonaPorDocumento("999")

	assert.NotNil(t, err)
	assert.Equal(t, "persona no encontrada", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestBuscarPersonaPorDocumento_ErrorBaseDeDatos(t *testing.T) {
	mockRepo := new(MockRepositorioBuscar)
	services.SetPersonaRepository(mockRepo)

	mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(models.Persona{}, errors.New("error de base de datos"))

	_, err := services.BuscarPersonaPorDocumento("123")

	assert.NotNil(t, err)
	assert.Equal(t, "error de base de datos", err.Error())
	mockRepo.AssertExpectations(t)
}
