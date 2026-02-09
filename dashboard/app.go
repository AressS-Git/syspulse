package main

import (
	"context"
	"fmt"
	"github.com/AressS-Git/syspulse/pkg/platform"
	"github.com/AressS-Git/syspulse/pkg/server"
)

// App será la estructura de nuestra app de Wails
type App struct {
	// Canal de comunicación entre código y Wails
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// startup es la función que Wails ejecuta una vez al abrir una ventana
// Por eso se guarda la conexión a la BD y el conexto (conexión código - wails) para luego usar ambas variables cómo queramos
func (appPointer *App) startup(ctx context.Context) {
    appPointer.ctx = ctx
    fmt.Println("Iniciando startup...")

    // Inicializar la conexión con la BD
    server.InitDB()
}

// GetStats es la función que React utilizará para sacar info de la BD
// Se devuelven slices vacíos si no se consiguen sacar datos de la BD
func (a *App) GetStats() []platform.SystemStats {
    // Si la BD no existe se informa por consola
    if server.DB == nil {
        fmt.Println("La base de datos aún no está conectada")
        return []platform.SystemStats{}
    }

    var systemData []platform.SystemStats
    
    // Hacer la query en la BD
    result := server.DB.Order("id desc").Limit(20).Find(&systemData)
    
    if result.Error != nil {
        fmt.Println("Error leyendo datos de la DB:", result.Error)
        return []platform.SystemStats{}
    }

    return systemData
}