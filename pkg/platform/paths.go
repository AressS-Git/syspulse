package platform

import (
	"fmt"
	"os"
    "path/filepath"
)

// Definir el nombre de la BD
const DBName = "syspulse.db"

// GetBDAbsolutePath se encargará de encontrar en el equipo la ruta de la BD syspulse.db
// Interesa que la ruta dónde se encuentre la BD sea en la ruta de configuración de usuario
// Aunque se elimine el programa toda la info quedará almacenada allí por si se vuelve a instalar
func GetBDAbsolutePath() (string, error) {
    // Obtener la ruta de ejecución del programa
    userConfigDir, err := os.UserConfigDir()
    if err != nil {
        fmt.Println("Error al encontrar la ruta del usuario:", err)
        return "", err
    }
    
    // Crear la carpeta SysPulse que contendrá la BD
    dbRoute := filepath.Join(userConfigDir, "SysPulse")

    // Crear la carpeta dentro de la configuración del usuario dónde se ubicará la BD
    if err := os.MkdirAll(dbRoute, 0700); err != nil {
        fmt.Println("Error al crear el directorio que contendrá la BD:", err)
        return "", nil
    }

    // Obtener la ruta completa: el directorio de la configuración del usuario + el nombre de la BD
    dbRouteWithBDName := filepath.Join(dbRoute, DBName)

    return dbRouteWithBDName, nil
}