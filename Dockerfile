# ---------- Etapa 1: Compilación ----------
    FROM golang:1.24.1 AS builder

    # Definir directorio de trabajo dentro del contenedor de compilación
    WORKDIR /app
    
    # Copiar archivos de dependencias y descargar módulos
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copiar el resto del código fuente
    COPY . .
    
    # Compilar el binario de la aplicación
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
    
    # ---------- Etapa 2: Imagen final optimizada ----------
    FROM debian:bullseye-slim
    
    # Crear usuario no-root para ejecución segura
    RUN useradd --create-home gouser
    
    # Definir directorio de trabajo para el usuario
    WORKDIR /home/gouser/app
    
    # Copiar binario compilado desde la etapa builder
    COPY --from=builder /app/app .
    
    # Establecer usuario no-root
    USER gouser
    
    # Comando por defecto al iniciar el contenedor
    CMD ["./app"]
    