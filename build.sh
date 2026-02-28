#!/bin/bash

# set -e hace que el script se detenga si cualquier comando falla
set -e

echo "--- 🛠️ Iniciando Limpieza de SysPulse ---"
# Limpiamos procesos y archivos (incluyendo los nuevos)
pkill sys-server || true
pkill sys-agent || true
rm -f sys-server sys-agent sys-agent-linux sys-agent-win.exe

echo "--- 🍎 Compilando para macOS (Originales) ---"
# Servidor para Mac
go build -o sys-server ./cmd/pulse-server/main.go
# Agente para Mac (el nombre que usabas siempre)
go build -o sys-agent ./cmd/pulse-agent/main.go

if [ -f "./sys-server" ] && [ -f "./sys-agent" ]; then
    echo "Binarios de macOS compilados con éxito."
else
    echo "Error: No se pudieron generar los binarios de macOS."
    exit 1
fi

echo "--- 🐧 Compilando Agente para LINUX ---"
GOOS=linux GOARCH=amd64 go build -o sys-agent-linux ./cmd/pulse-agent/main.go

echo "--- 🪟 Compilando Agente para WINDOWS ---"
GOOS=windows GOARCH=amd64 go build -o sys-agent-win.exe ./cmd/pulse-agent/main.go

echo "--- 📊 Compilando Dashboard (Wails) ---"
cd dashboard
wails build
cd ..

echo "--- ✅ Todo el sistema ha sido reconstruido ---"
echo "Archivos generados en la raíz:"
echo "  - sys-server          (Mac)"
echo "  - sys-agent           (Mac)"
echo "  - sys-agent-linux     (Linux)"
echo "  - sys-agent-win.exe   (Windows)"