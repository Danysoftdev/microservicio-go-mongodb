#!/bin/bash

# Script para construir imagen y levantar contenedores con Docker Compose

echo "🚀 Construyendo imagen Docker del microservicio..."
docker compose build --no-cache

echo "📦 Eliminando contenedores anteriores (si existen)..."
docker compose down

echo "🔄 Ejecutando los contenedores..."
docker compose up -d --build

echo "✅ Microservicio ejecutándose en http://localhost:8080"
