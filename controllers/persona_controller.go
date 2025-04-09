package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/danysoftdev/microservicio-go-mongodb/models"
	"github.com/danysoftdev/microservicio-go-mongodb/services"

	"github.com/gorilla/mux"
)

func CrearPersona(w http.ResponseWriter, r *http.Request) {
	var persona models.Persona

	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "El formato del cuerpo es inválido", http.StatusBadRequest)
		return
	}

	err = services.CrearPersona(persona)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Persona creada exitosamente"})
}

func ObtenerPersonas(w http.ResponseWriter, r *http.Request) {
	personas, err := services.ListarPersonas()
	if err != nil {
		http.Error(w, "Error al obtener personas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(personas)
}

func ObtenerPersonaPorDocumento(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documento := params["documento"]

	persona, err := services.BuscarPersonaPorDocumento(documento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(persona)
}

func ActualizarPersona(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documento := params["documento"]

	var persona models.Persona
	err := json.NewDecoder(r.Body).Decode(&persona)
	if err != nil {
		http.Error(w, "El formato del cuerpo es inválido", http.StatusBadRequest)
		return
	}

	err = services.ModificarPersona(documento, persona)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Persona actualizada exitosamente"})
}

func EliminarPersona(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documento := params["documento"]

	err := services.BorrarPersona(documento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Persona eliminada exitosamente"})
}
