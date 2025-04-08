package services_test

import (
	"errors"
	"testing"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/services"
	"github.com/danysoftdev/microservicio-go-mongodb/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestBorrarPersona_Exito(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
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
	mockRepo := new(mocks.MockPersonaRepo)
	services.Repo = mockRepo

	err := services.BorrarPersona(" ")

	assert.Error(t, err)
	assert.Equal(t, "el documento no puede estar vacío", err.Error())
}

func TestBorrarPersona_NoEncontrada(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", "000000").
		Return(models.Persona{}, mongo.ErrNoDocuments)

	err := services.BorrarPersona("000000")

	assert.Error(t, err)
	assert.Equal(t, "persona no encontrada", err.Error())
}

func TestBorrarPersona_ErrorEliminar(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", "987654").
		Return(models.Persona{Documento: "987654"}, nil)
	mockRepo.On("EliminarPersona", "987654").
		Return(errors.New("error al eliminar"))

	err := services.BorrarPersona("987654")

	assert.Error(t, err)
	assert.Equal(t, "error al eliminar", err.Error())
}
