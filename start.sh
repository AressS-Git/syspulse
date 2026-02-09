#!/bin/bash

echo "Iniciando Sistema SysPulse..."

# 1. Matar el agente y el servidor por si estaban abeiertos
pkill sys-server
pkill sys-agent

# 2. Arrancar el servidor en segundo plano
./sys-server &
echo "Servidor encendido"

# Esperar un poco antes de abrir el agente para darle tiempo al servidor
sleep 1

# 3. Arrancar el agente en segundo plano
./sys-agent &
echo "Agente activo"

# 4. Abrir el dashboard de Wails
open dashboard/build/bin/dashboard.app
echo "Dashboard abierto"