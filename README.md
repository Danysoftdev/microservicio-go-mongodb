# Microservicio CRUD en Golang con MongoDB

Este proyecto es un microservicio desarrollado en Golang que implementa un CRUD (Crear, Leer, Actualizar, Eliminar) sobre la entidad Persona. La aplicación utiliza MongoDB como base de datos y está diseñada para ser fácilmente desplegada mediante contenedores Docker y gestionada con Docker Compose. Además, incluye pruebas unitarias y de integración para garantizar su correcto funcionamiento.

## Características

- CRUD completo sobre la entidad Persona.
- Base de datos MongoDB.
- Contenerización con Docker.
- Orquestación de servicios con Docker Compose.
- Pruebas unitarias y de integración.

## Requisitos previos

- [Golang](https://golang.org/) instalado.
- [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/) instalados.
- MongoDB (opcional si se usa Docker Compose para levantar los servicios).

## Uso

El microservicio expone una API REST para interactuar con la entidad. A continuación, se describen los endpoints principales:

- **POST /persona**: Crear una nueva persona.
- **GET /personas**: Obtener todas las personas.
- **GET /persona/{id}**: Obtener una persona por su ID.
- **PUT /persona/{id}**: Actualizar una persona por su ID.
- **DELETE /persona/{id}**: Eliminar una persona por su ID.

## Tecnologías utilizadas

- Golang
- MongoDB
- Docker
- Docker Compose

