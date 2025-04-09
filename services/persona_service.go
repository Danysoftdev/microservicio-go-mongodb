package services

import (
	"errors"
	"strings"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

var Repo repositories.PersonaRepository

func SetPersonaRepository(r repositories.PersonaRepository) {
	Repo = r
}

func ValidarPersona(p models.Persona) error {
	if strings.TrimSpace(p.Documento) == "" {
		return errors.New("el documento no puede estar vacío")
	}
	if strings.TrimSpace(p.Nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	if strings.TrimSpace(p.Apellido) == "" {
		return errors.New("el apellido no puede estar vacío")
	}
	if p.Edad <= 0 {
		return errors.New("la edad debe ser un número entero mayor a 0")
	}
	if strings.TrimSpace(p.Correo) == "" || !strings.Contains(p.Correo, "@") {
		return errors.New("el correo es inválido")
	}
	if strings.TrimSpace(p.Telefono) == "" {
		return errors.New("el teléfono no puede estar vacío")
	}
	if strings.TrimSpace(p.Direccion) == "" {
		return errors.New("la dirección no puede estar vacía")
	}
	return nil
}

func CrearPersona(p models.Persona) error {
	if err := ValidarPersona(p); err != nil {
		return err
	}

	_, err := Repo.ObtenerPersonaPorDocumento(p.Documento)
	if err == nil {
		return errors.New("ya existe una persona con ese documento")
	}

	return Repo.InsertarPersona(p)
}

func ListarPersonas() ([]models.Persona, error) {
	return Repo.ObtenerPersonas()
}

func BuscarPersonaPorDocumento(doc string) (models.Persona, error) {
	if strings.TrimSpace(doc) == "" {
		return models.Persona{}, errors.New("el documento no puede estar vacío")
	}

	persona, err := Repo.ObtenerPersonaPorDocumento(doc)
	if err == mongo.ErrNoDocuments {
		return models.Persona{}, errors.New("persona no encontrada")
	}

	return persona, err
}

func ModificarPersona(documento string, p models.Persona) error {
	if strings.TrimSpace(documento) == "" {
		return errors.New("el documento no puede estar vacío")
	}

	if err := ValidarPersona(p); err != nil {
		return err
	}

	_, err := Repo.ObtenerPersonaPorDocumento(documento)
	if err == mongo.ErrNoDocuments {
		return errors.New("persona no encontrada")
	}

	if p.Documento != documento {
		return errors.New("no se puede modificar el documento de una persona")
	}

	return Repo.ActualizarPersona(documento, p)
}

func BorrarPersona(documento string) error {
	if strings.TrimSpace(documento) == "" {
		return errors.New("el documento no puede estar vacío")
	}

	_, err := Repo.ObtenerPersonaPorDocumento(documento)
	if err == mongo.ErrNoDocuments {
		return errors.New("persona no encontrada")
	}

	return Repo.EliminarPersona(documento)
}
