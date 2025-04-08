# ---------- Etapa 1: Compilación ----------
    FROM golang:1.24.1-alpine AS builder
    
    WORKDIR /app
    
    # Copiar dependencias y descargar módulos
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copiar el resto del código fuente
    COPY . .
    
    # Compilar el binario con soporte CGO (por defecto en Alpine)
    RUN go build -o app .
    
    # ---------- Etapa 2: Imagen final optimizada ----------
    FROM alpine:latest
    
    # Crear un usuario no-root para seguridad
    RUN adduser -D gouser
    
    # Definir directorio de trabajo
    WORKDIR /home/gouser/app
    
    # Copiar el binario desde la etapa anterior
    COPY --from=builder /app/app .
    
    # Cambiar al usuario seguro
    USER gouser
    
    # Ejecutar el binario
    CMD ["./app"]
    