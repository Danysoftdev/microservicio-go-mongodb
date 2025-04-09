package repositories

import (
	"context"
	"time"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Permite inyectar la colección desde fuera (ideal para pruebas)
func SetCollection(c *mongo.Collection) {
	collection = c
}

// InsertarPersona guarda una nueva persona en la base de datos
func InsertarPersona(persona models.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, persona)
	return err
}

// ObtenerPersonas devuelve una lista de todas las personas
func ObtenerPersonas() ([]models.Persona, error) {
	var personas []models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var persona models.Persona
		if err := cursor.Decode(&persona); err != nil {
			return nil, err
		}
		personas = append(personas, persona)
	}

	return personas, nil
}

// ObtenerPersonaPorDocumento busca una persona por su Documento
func ObtenerPersonaPorDocumento(documento string) (models.Persona, error) {
	var persona models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"documento": documento}).Decode(&persona)
	return persona, err
}

// ActualizarPersona actualiza los datos de una persona por Documento
func ActualizarPersona(documento string, persona models.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": persona,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"documento": documento}, update)
	return err
}

// EliminarPersona elimina una persona por su Documento
func EliminarPersona(documento string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"documento": documento})
	return err
}

type RealPersonaRepository struct{}

func (r RealPersonaRepository) InsertarPersona(p models.Persona) error {
	return InsertarPersona(p)
}

func (r RealPersonaRepository) ObtenerPersonas() ([]models.Persona, error) {
	return ObtenerPersonas()
}

func (r RealPersonaRepository) ObtenerPersonaPorDocumento(doc string) (models.Persona, error) {
	return ObtenerPersonaPorDocumento(doc)
}

func (r RealPersonaRepository) ActualizarPersona(doc string, p models.Persona) error {
	return ActualizarPersona(doc, p)
}

func (r RealPersonaRepository) EliminarPersona(doc string) error {
	return EliminarPersona(doc)
}