package main

import (
	"fmt"
	"log"
	"time"
	"github.com/AressS-Git/syspulse/pkg/agent"
)

func main() {
    // Definimos la URL donde está escuchando nuestro servidor
    serverURL := "http://192.168.1.16:8080/api/stats"

    fmt.Println("Iniciando Agente SysPulse...")

    // Bucle infinito para que el agente envíe estadísticas sin parar
    for {
        // Recolectar las estadísticas del equipo dónde está instalado el agente gracias a la funcion GetMetrics del archivo collector.go
        stats, err := agent.GetMetrics()
        if err != nil {
            log.Println("Error al recolectar métricas:", err)
            // Recojer métricas cada 5 segundos
            time.Sleep(5 * time.Second)
            continue
        }

        // Enviar las estadísticas al servidor mediante la función SendMetrics del archivo sender.go
        err = agent.SendMetrics(serverURL, stats)
        if err != nil {
            log.Println("Error al enviar métricas al servidor:", err)
        } else {
            // Si todo va bien, se muestra un mensaje por consola de las principales métricas (CPU y RAM)
            fmt.Printf("Métricas enviadas al servidor -> CPU = %.1f%%, RAM = %.1f%%", stats.CpuUsage, stats.RamUsage) // Se muestran sólo el primer decimal
        }

        // El agente enviará datos cada 10 segundos para no saturar el servidor
        time.Sleep(10 * time.Second)
    }
}