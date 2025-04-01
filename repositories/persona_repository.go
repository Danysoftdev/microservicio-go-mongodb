package repositories

import (
	"context"
	"time"

	"github.com/danysoftdev/microservicio-go-mongodb/config"
	"github.com/danysoftdev/microservicio-go-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertarPersona guarda una nueva persona en la base de datos
func InsertarPersona(persona models.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := config.Collection.InsertOne(ctx, persona)
	return err
}

// ObtenerPersonas devuelve una lista de todas las personas
func ObtenerPersonas() ([]models.Persona, error) {
	var personas []models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.Collection.Find(ctx, bson.M{})
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

// ObtenerPersonaPorID busca una persona por su ID
func ObtenerPersonaPorID(id string) (models.Persona, error) {
	var persona models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	err := config.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&persona)
	return persona, err
}

// ActualizarPersona actualiza los datos de una persona por ID
func ActualizarPersona(id string, persona models.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": persona,
	}

	_, err := config.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// EliminarPersona elimina una persona por su ID
func EliminarPersona(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := config.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
