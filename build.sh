#!/bin/bash

# set -e hace que el script se detenga si cualquier comando falla
set -e

echo "--- 🛠️ Iniciando Limpieza de SysPulse ---"
# Limpiamos procesos y archivos
pkill sys-server || true
pkill sys-agent || true
rm -f sys-server sys-agent sys-agent-linux syspulse-completo

echo "--- 🍎 Compilando para macOS ---"
# Servidor para Mac
go build -o sys-server ./cmd/pulse-server/main.go
# Agente para Mac
go build -o sys-agent ./cmd/pulse-agent/main.go

if [ -f "./sys-server" ] && [ -f "./sys-agent" ]; then
    echo "Binarios de macOS compilados con éxito."
else
    echo "Error: No se pudieron generar los binarios de macOS."
    exit 1
fi

echo "--- 🐧 Compilando Agente para LINUX ARM64 ---"
GOOS=linux GOARCH=arm64 go build -o sys-agent-linux-arm64 ./cmd/pulse-agent/main.go


echo "--- 📊 Compilando Dashboard (Wails) ---"
cd dashboard
wails build
cd ..

echo "--- 📦 EMPAQUETANDO TODO EN UN SOLO ARCHIVO ---"
# Limpiamos posibles restos
rm -f SysPulse.command

# Creamos un comprimido temporal con los binarios y el dashboard
tar -czf proyecto_temp.tar.gz sys-server sys-agent dashboard/build/bin/dashboard.app

# Creamos el script lanzador auto-extraíble
cat <<'EOF' > SysPulse.command
#!/bin/bash
# SysPulse Self-Extracting Bundle

# Crear un directorio temporal para la ejecución
TMP_DIR=$(mktemp -d /tmp/syspulse.XXXXXX)

# Función para limpiar al cerrar
cleanup() {
    echo "Stopping services..."
    pkill -9 -f "sys-server" || true
    pkill -9 -f "sys-agent" || true
    rm -rf "$TMP_DIR"
    exit
}
trap cleanup SIGINT SIGTERM EXIT

# Extraer el contenido comprimido interno
echo "🚀 Iniciando SysPulse... (Preparando entorno)"
sed '1,/^#PAYLOAD#$/d' "$0" | tar -xz -C "$TMP_DIR"

# Entrar al directorio y lanzar todo
cd "$TMP_DIR"

# 1. Servidor
./sys-server > /dev/null 2>&1 &
# 2. Agente
./sys-agent > /dev/null 2>&1 &

echo "✅ Servicios activos. Abriendo Dashboard..."
# 3. Abrir la interfaz y esperar a que se cierre
open -W dashboard/build/bin/dashboard.app

# El script termina cuando se cierra la ventana
cleanup
exit 0
#PAYLOAD#
EOF

# Añadir los datos comprimidos al final del archivo
cat proyecto_temp.tar.gz >> SysPulse.command
chmod +x SysPulse.command
rm proyecto_temp.tar.gz

echo "--- ✅ Reconstrucción completa ---"
echo "Archivos generados:"
echo "  - sys-server       (Servidor)"
echo "  - sys-agent        (Agente)"
echo "  - SysPulse.command (📦 EJECUTABLE ÚNICO PARA ENTREGA)"