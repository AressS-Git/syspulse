package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
    // Instanciar la estructura de la app
    app := NewApp()

    // Ejecutar wails
    err := wails.Run(&options.App{
        Title:  "SysPulse - System Monitor", // Nombre más descriptivo
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
        
        OnStartup: app.startup, 
        
        Bind: []interface{}{
            app,
        },
    })

    if err != nil {
        fmt.Printf("Error fatal al iniciar SysPulse: %v", err)
    }
}