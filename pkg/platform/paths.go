package platform

import (
	"fmt"
	"os"
    "path/filepath"
)

// Definir el nombre de la BD
const dbName = "syspulse.db"

// GetBDAbsolutePath se encargará de encontrar en el equipo la ruta de la BD syspulse.db
// Esto permite que si la BD no se tenga que encontrar si o si en un directorio concreto
func GetBDAbsolutePath() (string, error) {
    // Obtener la ruta de ejecución del programa
    actualExecutionFileRoute, err := os.Executable()
    if err != nil {
        fmt.Println("Error al encontrar la ruta del ejecutable:", err)
        return "", err
    }

    // Obtener el directorio del ejecutable
    actualExecutionDir := filepath.Dir(actualExecutionFileRoute)

    // Unir las rutas para crear la ruta dónde está la BD
    dbRoute := filepath.Join(actualExecutionDir, dbName)

    return dbRoute, nil
}