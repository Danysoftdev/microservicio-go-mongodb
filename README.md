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

- **POST /crear-personas**: Crear una nueva persona.
- **GET /buscar-personas**: Obtener todas las personas.
- **GET /buscar-personas/{documento}**: Obtener una persona por su documento.
- **PUT /actualizar-personas/{documento}**: Actualizar una persona por su documento.
- **DELETE /eliminar-persona/{documento}**: Eliminar una persona por su documento.

## Tecnologías utilizadas

- Golang
- MongoDB
- Docker
- Docker Compose

