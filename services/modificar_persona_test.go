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

// MockPersonaModificar implementa la interfaz PersonaRepository para las pruebas
type MockPersonaModificar struct {
	mock.Mock
}

func (m *MockPersonaModificar) InsertarPersona(p models.Persona) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPersonaModificar) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}

func (m *MockPersonaModificar) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockPersonaModificar) ActualizarPersona(doc string, p models.Persona) error {
	args := m.Called(doc, p)
	return args.Error(0)
}

func (m *MockPersonaModificar) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}

func TestModificarPersona(t *testing.T) {
	mockRepo := new(MockPersonaModificar)
	services.SetPersonaRepository(mockRepo)

	personaValida := models.Persona{
		Documento: "123",
		Nombre:    "Laura",
		Apellido:  "Gomez",
		Edad:      25,
		Correo:    "laura@example.com",
		Telefono:  "555-1234",
		Direccion: "Calle Falsa 123",
	}

	t.Run("Debe modificar una persona exitosamente", func(t *testing.T) {
		mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(personaValida, nil)
		mockRepo.On("ActualizarPersona", "123", personaValida).Return(nil)

		err := services.ModificarPersona("123", personaValida)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Debe fallar si el documento está vacío", func(t *testing.T) {
		err := services.ModificarPersona("", personaValida)
		assert.EqualError(t, err, "el documento no puede estar vacío")
	})

	t.Run("Debe fallar si la validación de datos es incorrecta", func(t *testing.T) {
		invalida := personaValida
		invalida.Nombre = ""

		err := services.ModificarPersona("123", invalida)
		assert.EqualError(t, err, "el nombre no puede estar vacío")
	})

	t.Run("Debe fallar si se intenta cambiar el documento", func(t *testing.T) {
		nueva := personaValida
		nueva.Documento = "456"

		err := services.ModificarPersona("123", nueva)
		assert.EqualError(t, err, "no se puede modificar el documento de una persona")
	})

    t.Run("Debe fallar si la persona no existe", func(t *testing.T) {
		mockRepo := new(MockPersonaModificar)
		services.SetPersonaRepository(mockRepo)

		mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(models.Persona{}, mongo.ErrNoDocuments)

		err := services.ModificarPersona("123", personaValida)
		assert.EqualError(t, err, "persona no encontrada")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Debe retornar error si falla la actualización", func(t *testing.T) {
		mockRepo := new(MockPersonaModificar)
		services.SetPersonaRepository(mockRepo)

		mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(personaValida, nil)
		mockRepo.On("ActualizarPersona", "123", personaValida).Return(errors.New("error al actualizar"))

		err := services.ModificarPersona("123", personaValida)
		assert.EqualError(t, err, "error al actualizar")
		mockRepo.AssertExpectations(t)
	})

}
