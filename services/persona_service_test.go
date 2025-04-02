package services

import (
	"testing"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/stretchr/testify/assert"
)

func TestValidarPersona_Valido(t *testing.T) {
	persona := models.Persona{
		Documento: "123",
		Nombre:    "Angie",
		Apellido:  "Diaz",
		Edad:      25,
		Correo:    "angie@email.com",
		Telefono:  "1234567890",
		Direccion: "Calle 123",
	}

	err := ValidarPersona(persona)
	assert.NoError(t, err)
}

func TestValidarPersona_CamposVacios(t *testing.T) {
	pruebas := []struct {
		nombre       string
		persona      models.Persona
		mensajeError string
	}{
		{
			"Documento vacío",
			models.Persona{Documento: "", Nombre: "A", Apellido: "B", Edad: 20, Correo: "a@a.com", Telefono: "123", Direccion: "Dir"},
			"el documento no puede estar vacío",
		},
		{
			"Nombre vacío",
			models.Persona{Documento: "123", Nombre: "", Apellido: "B", Edad: 20, Correo: "a@a.com", Telefono: "123", Direccion: "Dir"},
			"el nombre no puede estar vacío",
		},
		{
			"Apellido vacío",
			models.Persona{Documento: "123", Nombre: "A", Apellido: "", Edad: 20, Correo: "a@a.com", Telefono: "123", Direccion: "Dir"},
			"el apellido no puede estar vacío",
		},
		{
			"Edad inválida",
			models.Persona{Documento: "123", Nombre: "A", Apellido: "B", Edad: 0, Correo: "a@a.com", Telefono: "123", Direccion: "Dir"},
			"la edad debe ser un número entero mayor a 0",
		},
		{
			"Correo inválido",
			models.Persona{Documento: "123", Nombre: "A", Apellido: "B", Edad: 20, Correo: "correo.com", Telefono: "123", Direccion: "Dir"},
			"el correo es inválido",
		},
		{
			"Teléfono vacío",
			models.Persona{Documento: "123", Nombre: "A", Apellido: "B", Edad: 20, Correo: "a@a.com", Telefono: "", Direccion: "Dir"},
			"el teléfono no puede estar vacío",
		},
		{
			"Dirección vacía",
			models.Persona{Documento: "123", Nombre: "A", Apellido: "B", Edad: 20, Correo: "a@a.com", Telefono: "123", Direccion: ""},
			"la dirección no puede estar vacía",
		},
	}

	for _, tt := range pruebas {
		t.Run(tt.nombre, func(t *testing.T) {
			err := ValidarPersona(tt.persona)
			assert.EqualError(t, err, tt.mensajeError)
		})
	}
}

func TestBuscarPersonaPorDocumento_ErrorDocumentoVacio(t *testing.T) {
	_, err := BuscarPersonaPorDocumento("")
	assert.EqualError(t, err, "el documento no puede estar vacío")
}

func TestModificarPersona_DocumentoNoEditable(t *testing.T) {
	persona := models.Persona{
		Documento: "456",
		Nombre:    "Nuevo",
		Apellido:  "Apellido",
		Edad:      30,
		Correo:    "nuevo@email.com",
		Telefono:  "321",
		Direccion: "Calle nueva",
	}

	err := ModificarPersona("123", persona)
	assert.EqualError(t, err, "no se puede modificar el documento de una persona")
}

func TestBorrarPersona_DocumentoVacio(t *testing.T) {
	err := BorrarPersona("")
	assert.EqualError(t, err, "el documento no puede estar vacío")
}

/*
func TestBuscarPersonaPorDocumento_NoExiste(t *testing.T) {
    _, err := BuscarPersonaPorDocumento("no-existe-999")
    assert.EqualError(t, err, "persona no encontrada")
}

func TestModificarPersona_NoExiste(t *testing.T) {
    persona := models.Persona{
        Documento: "no-existe-999",
        Nombre:    "Cambio",
        Apellido:  "Apellido",
        Edad:      30,
        Correo:    "nuevo@email.com",
        Telefono:  "321",
        Direccion: "Nueva calle",
    }

    err := ModificarPersona("no-existe-999", persona)
    assert.EqualError(t, err, "persona no encontrada")
}

func TestBorrarPersona_NoExiste(t *testing.T) {
    err := BorrarPersona("no-existe-999")
    assert.EqualError(t, err, "persona no encontrada")
}
*/
