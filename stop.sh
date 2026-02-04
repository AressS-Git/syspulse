#!/bin/bash

echo "Deteniendo todo el sistema SysPulse..."

# Matamos los procesos
pkill sys-server
pkill sys-agent

# Opcional: También podemos matar la app gráfica si quieres cerrarla de golpe
pkill dashboard

echo "Sistema apagado correctamente."