package server

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/AressS-Git/syspulse/pkg/platform"
)

// Variable global que permite innteractuar con la BD desde otros paquetes (principalmente desde la app de Wails)
var DB *gorm.DB

// Mensaje de error por defecto si falla la conexión a la BD
const conexionError string = "Error fatal, no se pudo abrir una conexión a la BD: "

// ConnectDB abre el canal con la BD
func ConnectDB() *gorm.DB {
    dbPath, err := platform.GetBDAbsolutePath()
    if err != nil {
        panic(conexionError + err.Error())
    }
    
    // Abrir la conexión con la BD
    dbConnexion, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        panic(conexionError + err.Error())
    }
    
    return dbConnexion // Devolver el puntero de la BD
}

// InitDB comprueba si la BD tiene las tablas necesarias para guardar los datos de los agentes
func InitDB() {
    // Obtener el puntero de conexión a la BD usando la función ConnectDB
    dbConn := ConnectDB()

    // Ejecutar las migraciones en base a la estrutura del struct SystemStats
    err := dbConn.AutoMigrate(&platform.SystemStats{})
    if err != nil {
        panic("Fallo en la migración de la base de datos: " + err.Error())
    }

    // Guardar la conexión con la BD en una variable global
    DB = dbConn
}