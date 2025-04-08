package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/services"
)

// MockPersonaRepository es un mock de la interfaz PersonaRepository
type MockPersonaRepository struct {
	mock.Mock
}

func (m *MockPersonaRepository) InsertarPersona(p models.Persona) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockPersonaRepository) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}

func (m *MockPersonaRepository) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	args := m.Called(doc)
	return args.Get(0).(models.Persona), args.Error(1)
}

func (m *MockPersonaRepository) ActualizarPersona(doc string, p models.Persona) error {
	args := m.Called(doc, p)
	return args.Error(0)
}

func (m *MockPersonaRepository) EliminarPersona(doc string) error {
	args := m.Called(doc)
	return args.Error(0)
}

// TestCrearPersonaExitosa prueba la creación exitosa de una persona
func TestCrearPersonaExitosa(t *testing.T) {
	mockRepo := new(MockPersonaRepository)
	services.Repo = mockRepo

	persona := models.Persona{
		Documento: "123",
		Nombre:    "Ana",
		Apellido:  "Díaz",
		Edad:      25,
		Correo:    "ana@example.com",
		Telefono:  "1234567",
		Direccion: "Calle Falsa 123",
	}

	mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(models.Persona{}, errors.New("not found"))
	mockRepo.On("InsertarPersona", persona).Return(nil)

	err := services.CrearPersona(persona)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestCrearPersonaYaExiste prueba cuando el documento ya está registrado
func TestCrearPersonaYaExiste(t *testing.T) {
	mockRepo := new(MockPersonaRepository)
	services.Repo = mockRepo

	persona := models.Persona{
		Documento: "123",
		Nombre:    "Ana",
		Apellido:  "Díaz",
		Edad:      25,
		Correo:    "ana@example.com",
		Telefono:  "1234567",
		Direccion: "Calle Falsa 123",
	}

	mockRepo.On("ObtenerPersonaPorDocumento", "123").Return(persona, nil)

	err := services.CrearPersona(persona)
	assert.EqualError(t, err, "ya existe una persona con ese documento")
}

// TestCrearPersonaConDatosInvalidos prueba todos los errores de validación
func TestCrearPersonaConDatosInvalidos(t *testing.T) {
	mockRepo := new(MockPersonaRepository)
	services.Repo = mockRepo

	mockRepo.On("ObtenerPersonaPorDocumento", mock.Anything).Return(models.Persona{}, errors.New("not found"))

	casos := []struct {
		nombre         string
		persona        models.Persona
		errorEsperado  string
	}{
		{"Documento vacío", models.Persona{Nombre: "Ana", Apellido: "Díaz", Edad: 25, Correo: "ana@example.com", Telefono: "123", Direccion: "Calle"}, "el documento no puede estar vacío"},
		{"Nombre vacío", models.Persona{Documento: "123", Apellido: "Díaz", Edad: 25, Correo: "ana@example.com", Telefono: "123", Direccion: "Calle"}, "el nombre no puede estar vacío"},
		{"Apellido vacío", models.Persona{Documento: "123", Nombre: "Ana", Edad: 25, Correo: "ana@example.com", Telefono: "123", Direccion: "Calle"}, "el apellido no puede estar vacío"},
		{"Edad inválida", models.Persona{Documento: "123", Nombre: "Ana", Apellido: "Díaz", Edad: 0, Correo: "ana@example.com", Telefono: "123", Direccion: "Calle"}, "la edad debe ser un número entero mayor a 0"},
		{"Correo inválido", models.Persona{Documento: "123", Nombre: "Ana", Apellido: "Díaz", Edad: 25, Correo: "anaexample.com", Telefono: "123", Direccion: "Calle"}, "el correo es inválido"},
		{"Teléfono vacío", models.Persona{Documento: "123", Nombre: "Ana", Apellido: "Díaz", Edad: 25, Correo: "ana@example.com", Telefono: "", Direccion: "Calle"}, "el teléfono no puede estar vacío"},
		{"Dirección vacía", models.Persona{Documento: "123", Nombre: "Ana", Apellido: "Díaz", Edad: 25, Correo: "ana@example.com", Telefono: "123", Direccion: ""}, "la dirección no puede estar vacía"},
	}

	for _, tt := range casos {
		t.Run(tt.nombre, func(t *testing.T) {
			err := services.CrearPersona(tt.persona)
			assert.EqualError(t, err, tt.errorEsperado)
		})
	}
}
