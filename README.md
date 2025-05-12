# Microservicio CRUD en Golang con MongoDB

Este proyecto es un microservicio desarrollado en Golang que implementa un CRUD (Crear, Leer, Actualizar, Eliminar) sobre la entidad Persona. La aplicación utiliza MongoDB como base de datos y está diseñada para ser fácilmente desplegada mediante contenedores Docker y gestionada con Docker Compose. Además, incluye pruebas unitarias y de integración automatizadas mediante GitHub Actions para garantizar su correcto funcionamiento y la calidad del código.

## Características

- CRUD completo sobre la entidad Persona.
- Base de datos MongoDB.
- Contenerización con Docker.
- Orquestación de servicios con Docker Compose.
- Pruebas unitarias y de integración automatizadas con GitHub Actions.
- Publicación automática de imágenes Docker a GitHub Container Registry (GHCR) y Docker Hub.
- Creación automática de GitHub Releases al crear tags con formato `v*.*.*`.
- Análisis de seguridad de la imagen Docker con Trivy.

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

## Integración Continua

Este proyecto utiliza GitHub Actions para automatizar el proceso de build, testeo, análisis de seguridad y publicación de la imagen Docker. El workflow se activa en los siguientes eventos:

- **`push` en la rama `main` (o tu rama principal):** Ejecuta las pruebas unitarias y de integración, construye la imagen Docker, la analiza en busca de vulnerabilidades con Trivy y, si las pruebas pasan y no se encuentran vulnerabilidades críticas o altas, publica la imagen en GitHub Container Registry y Docker Hub.

- **`push` de tags con el patrón `v*.*.*` (ej: `v1.0.0`):** Similar al push en la rama principal, pero además crea automáticamente un GitHub Release con las notas generadas a partir de los commits.

El workflow `build-and-push.yml` realiza los siguientes pasos:

1.  **Checkout del repositorio:** Obtiene el código fuente.

2.  **Configuración de Go:** Establece la versión de Go utilizada en el proyecto.

3.  **Ejecución de pruebas unitarias:** Ejecuta todas las pruebas unitarias excluyendo las de integración.

4.  **Ejecución de pruebas de integración (Testcontainers):** Ejecuta las pruebas de integración que utilizan Testcontainers.

5.  **Configuración de Docker Buildx:** Prepara el entorno para la construcción de imágenes Docker.

6.  **Login en GitHub Container Registry:** Autentica con GHCR utilizando el token de GitHub.

7.  **Login en Docker Hub:** Autentica con Docker Hub utilizando las credenciales almacenadas como secretos.

8.  **Creación del archivo `.env`:** Crea un archivo `.env` con las variables de entorno necesarias para Docker Compose, utilizando secretos de GitHub.

9.  **Ejecución de pruebas de integración con Docker Compose:** Levanta los servicios definidos en `docker-compose.yml` con el perfil `test`, ejecuta pruebas y luego los detiene.

10. **Extracción de metadatos para Docker:** Genera tags y labels para la imagen Docker basados en la rama y las etiquetas Git.

11. **Construcción de la imagen Docker (local):** Construye la imagen Docker sin publicarla inmediatamente.

12. **Análisis de la imagen Docker con Trivy:** Escanea la imagen en busca de vulnerabilidades de seguridad críticas y altas. El workflow fallará si se encuentran vulnerabilidades con estos niveles de severidad.

13. **Publicación de la imagen Docker:** Si el análisis de Trivy no encuentra problemas, la imagen Docker se publica en GitHub Container Registry y Docker Hub.

14. **Creación de GitHub Release:** Si el evento que activó el workflow fue la creación de un tag (con el patrón `v*.*.*`), se crea automáticamente un nuevo GitHub Release con las notas generadas.

## Tecnologías utilizadas

- Golang
- MongoDB
- Docker
- Docker Compose
- GitHub Actions
- Trivy