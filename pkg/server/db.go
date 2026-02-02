package server

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/AressS-Git/syspulse/pkg/platform"
)

func InitDB() (*gorm.DB, error) {
    // Se crea la conexión de gorm con la BD
    db, err := gorm.Open(sqlite.Open("syspulse.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Gorm revisa la estructura que le pases (en este caso SystemStats) y si no existe en la BD crea una tabla
    if err := db.AutoMigrate(&platform.SystemStats{}); err != nil {
        return nil, err
    }

    return db, nil
}