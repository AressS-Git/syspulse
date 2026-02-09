#!/bin/bash

# set -e hace que el script se detenga si cualquier comando falla
set -e

echo "--- 🛠️ Iniciando Limpieza de SysPulse ---"
# Matar procesos antiguos para liberar los archivos y los puertos
pkill sys-server || true
pkill sys-agent || true
rm -f sys-server sys-agent

echo "--- Compilando Servidor ---"
# Compilamos el servidor y verificamos que el archivo existe
go build -o sys-server ./cmd/pulse-server/main.go
if [ -f "./sys-server" ]; then
    echo "Servidor compilado con éxito."
else
    echo "Error: No se pudo generar el binario del servidor."
    exit 1
fi

echo "--- Compilando Agente ---"
go build -o sys-agent ./cmd/pulse-agent/main.go
if [ -f "./sys-agent" ]; then
    echo "Agente compilado con éxito."
else
    echo "Error: No se pudo generar el binario del agente."
    exit 1
fi

echo "--- Compilando Dashboard (Wails) ---"
# Entramos a la carpeta, construimos y volvemos a la raíz
cd dashboard
wails build
cd ..

echo "--- Todo el sistema ha sido reconstruido con éxito ---"