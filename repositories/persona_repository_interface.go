package repositories

import "github.com/danysoftdev/microservicio-go-mongodb/models"

type PersonaRepository interface {
	InsertarPersona(persona models.Persona) error
	ObtenerPersonas() ([]models.Persona, error)
	ObtenerPersonaPorDocumento(documento string) (models.Persona, error)
	ActualizarPersona(documento string, persona models.Persona) error
	EliminarPersona(documento string) error
}
