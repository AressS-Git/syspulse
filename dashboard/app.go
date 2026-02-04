package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"github.com/AressS-Git/syspulse/pkg/platform"
	"github.com/AressS-Git/syspulse/pkg/server"
)

// App será la estructura de nuestra app de Wails
type App struct {
	// Canal de comunicación entre código y Wails
	ctx context.Context
	// Base de datos que usará Wails
	db *gorm.DB
}

func NewApp() *App {
	return &App{}
}

// startup es la función que Wails ejecuta una vez al abrir una ventana
// de ahí que se guarde la conexión a la BD y el conexto (conexión código - wails) para luego usar ambas variables cómo queramos
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    fmt.Println("INICIANDO STARTUP...")

    // ESTA ES TU RUTA REAL (Extraída de tu log)
    // Apunta al archivo syspulse.db que está en la carpeta superior a dashboard
    fullPath := "/Users/sergiogomezsantos/Desktop/syspulse/syspulse.db"
    
    fmt.Println("Intentando conectar a:", fullPath)

    conn, err := server.InitDB(fullPath)
    if err != nil {
        // Usamos panic para que el error sea muy visible si falla
        panic(fmt.Sprintf("ERROR FATAL EN INITDB: %v", err))
    }
    
    a.db = conn
    fmt.Println("CONEXIÓN ÉXITOSA - ¡Base de datos cargada!")
}

// GetStats es la función que React utilizará para sacar info de la BD
func (a *App) GetStats() []platform.SystemStats {
    // Si la BD no existe se informa por consola
    if a.db == nil {
        fmt.Println("La base de datos aún no está conectada.")
        return []platform.SystemStats{}
    }

    var systemData []platform.SystemStats
    
    // Hacer la query en la BD
    result := a.db.Order("id desc").Limit(20).Find(&systemData)
    
    if result.Error != nil {
        fmt.Println("Error leyendo datos:", result.Error)
    }

    return systemData
}