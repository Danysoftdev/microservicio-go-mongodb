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

// Reutilizamos el mock ya definido
type MockRepositorioBorrar struct {
	mock.Mock
}

func (m *MockRepositorioBorrar) InsertarPersona(p models.Persona) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockRepositorioBorrar) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}

func (m *MockRepositorioBorrar) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockRepositorioBorrar) ActualizarPersona(doc string, p models.Persona) error {
	args := m.Called(doc, p)
	return args.Error(0)
}

func (m *MockRepositorioBorrar) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}

func TestBorrarPersona_Exito(t *testing.T) {
	mockRepo := new(MockRepositorioBorrar)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", "123456").
		Return(models.Persona{Documento: "123456"}, nil)
	mockRepo.On("EliminarPersona", "123456").
		Return(nil)

	err := services.BorrarPersona("123456")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBorrarPersona_DocumentoVacio(t *testing.T) {
	mockRepo := new(MockRepositorioBorrar)
	services.Repo = mockRepo

	err := services.BorrarPersona(" ")

	assert.Error(t, err)
	assert.Equal(t, "el documento no puede estar vacío", err.Error())
}

func TestBorrarPersona_NoEncontrada(t *testing.T) {
	mockRepo := new(MockRepositorioBorrar)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", "000000").
		Return(models.Persona{}, mongo.ErrNoDocuments)

	err := services.BorrarPersona("000000")

	assert.Error(t, err)
	assert.Equal(t, "persona no encontrada", err.Error())
}

func TestBorrarPersona_ErrorEliminar(t *testing.T) {
	mockRepo := new(MockRepositorioBorrar)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", "987654").
		Return(models.Persona{Documento: "987654"}, nil)
	mockRepo.On("EliminarPersona", "987654").
		Return(errors.New("error al eliminar"))

	err := services.BorrarPersona("987654")

	assert.Error(t, err)
	assert.Equal(t, "error al eliminar", err.Error())
}
