#!/bin/bash

# Ruta exacta de la base de datos en macOS según os.UserConfigDir() de Go
DB_PATH="$HOME/Library/Application Support/SysPulse/syspulse.db"

echo "🔍 Buscando la base de datos en: $DB_PATH"

# Comprobar si el archivo de la BD realmente existe
if [ ! -f "$DB_PATH" ]; then
    echo "❌ Error: No se encontró la base de datos. Quizás aún no se ha creado."
    exit 1
fi

echo "🧹 Limpiando datos de las tablas (manteniendo la estructura)..."

# Ejecutar comandos SQLite para vaciar las tablas
# El orden es importante por las claves foráneas: primero borramos hijas, luego padre
sqlite3 "$DB_PATH" <<EOF
DELETE FROM alerts;
DELETE FROM system_stats;
DELETE FROM devices;

-- Reiniciar los contadores de ID (auto-increment) para que empiecen desde 1
DELETE FROM sqlite_sequence WHERE name='alerts';
DELETE FROM sqlite_sequence WHERE name='system_stats';
DELETE FROM sqlite_sequence WHERE name='devices';

-- Optimizar y liberar el espacio en el disco
VACUUM;
EOF

echo "✅ ¡Base de datos limpiada con éxito!"