#!/bin/bash

echo "Iniciando Sistema SysPulse..."

# 1. Matamos versiones antiguas por si acaso se quedaron colgadas
pkill sys-server
pkill sys-agent

# 2. Arrancamos el Servidor en SEGUNDO PLANO (el símbolo & es la clave)
#    Así la terminal no se queda bloqueada.
./sys-server &
echo "Servidor activo"

# Esperamos 1 segundo para dar tiempo al servidor
sleep 1

# 3. Arrancamos el Agente en SEGUNDO PLANO
./sys-agent &
echo "Agente activo"

# 4. Abrimos la App gráfica (MacOS usa el comando 'open')
open dashboard/build/bin/dashboard.app
echo "Dashboard abierto"