package services_test

import (
	"errors"
	"testing"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
    "github.com/danysoftdev/microservicio-go-mongodb/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepositorio implementa la interfaz PersonaRepository
type MockRepositorio struct {
	mock.Mock
}

func (m *MockRepositorio) InsertarPersona(p models.Persona) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockRepositorio) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}

func (m *MockRepositorio) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockRepositorio) ActualizarPersona(doc string, p models.Persona) error {
	args := m.Called(doc, p)
	return args.Error(0)
}

func (m *MockRepositorio) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}

func TestListarPersonas_Success(t *testing.T) {
	mockRepo := new(MockRepositorio)
    services.Repo = mockRepo

	personasMock := []models.Persona{
		{Documento: "123", Nombre: "Ana", Apellido: "Díaz"},
		{Documento: "456", Nombre: "Luis", Apellido: "Pérez"},
	}

	mockRepo.On("ObtenerPersonas").Return(personasMock, nil)
	services.SetPersonaRepository(mockRepo) // 👈 Aquí se inyecta el mock

	personas, err := services.ListarPersonas()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(personas))
	assert.Equal(t, "Ana", personas[0].Nombre)
	mockRepo.AssertExpectations(t)
}

func TestListarPersonas_Error(t *testing.T) {
	mockRepo := new(MockRepositorio)
    services.Repo = mockRepo

	mockRepo.On("ObtenerPersonas").Return([]models.Persona(nil), errors.New("fallo al obtener"))
	services.SetPersonaRepository(mockRepo)

	personas, err := services.ListarPersonas()

	assert.Error(t, err)
	assert.Nil(t, personas)
	mockRepo.AssertExpectations(t)
}
