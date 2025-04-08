package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danysoftdev/microservicio-go-mongodb/controllers"
	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/services"
	"github.com/danysoftdev/microservicio-go-mongodb/tests/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestEliminarPersonaController_Success(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.SetPersonaRepository(mockRepo)

	doc := "123"
	mockRepo.On("ObtenerPersonaPorDocumento", doc).Return(models.Persona{Documento: doc}, nil)
	mockRepo.On("EliminarPersona", doc).Return(nil)

	req := httptest.NewRequest("DELETE", "/personas/"+doc, nil)
	req = mux.SetURLVars(req, map[string]string{"documento": doc})
	rr := httptest.NewRecorder()

	controllers.EliminarPersona(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Persona eliminada exitosamente", response["mensaje"])

	mockRepo.AssertExpectations(t)
}

func TestEliminarPersonaController_DocumentoVacio(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/personas/", nil)
	req = mux.SetURLVars(req, map[string]string{"documento": ""})
	rr := httptest.NewRecorder()

	controllers.EliminarPersona(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "el documento no puede estar vacío")
}


func TestEliminarPersonaController_ErrorEliminar(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.SetPersonaRepository(mockRepo)

	doc := "456"
	mockRepo.On("ObtenerPersonaPorDocumento", doc).Return(models.Persona{Documento: doc}, nil)
	mockRepo.On("EliminarPersona", doc).Return(errors.New("fallo al eliminar"))

	req := httptest.NewRequest("DELETE", "/personas/"+doc, nil)
	req = mux.SetURLVars(req, map[string]string{"documento": doc})
	rr := httptest.NewRecorder()

	controllers.EliminarPersona(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "fallo al eliminar")

	mockRepo.AssertExpectations(t)
}
