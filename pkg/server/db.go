package server

import (
    "fmt"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/AressS-Git/syspulse/pkg/platform"
)

func InitDB(connectionRoute string) (*gorm.DB, error) {
    // Se crea la conexión de gorm con la BD
    db, err := gorm.Open(sqlite.Open(connectionRoute), &gorm.Config{})
    if err != nil {
        fmt.Println("Error al abrir Gorm:", err)
        return nil, err
    }

    // Gorm revisa la estructura que le pases (en este caso SystemStats) y si no existe en la BD crea una tabla
    if err := db.AutoMigrate(&platform.SystemStats{}); err != nil {
        fmt.Println("Error en AutoMigrate:", err)
        return nil, err
    }

    return db, nil
}