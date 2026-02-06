#!/bin/bash

echo "Deteniendo todo el sistema SysPulse..."

# Matar el agente y el servidor
pkill sys-server
pkill sys-agent

# Matar el dashboard de Wails
pkill dashboard

echo "Sistema apagado correctamente"