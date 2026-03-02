package server

import (
    "github.com/AressS-Git/syspulse/pkg/platform"
    "fmt"
    "time"
)

// Constantes que servirán cómo umbrales para decidir si algo es una alerta o no lo es
const (
    CPU_THRESHOLD float64 = 95.0
    RAM_THRESHOLD float64 = 95.0
    DISK_THRESHOLD float64 = 90.0
)

// CreateAlert creará una alerta si las estadísticas de la CPU, RAM y Disco superan los ciertos límites
func CreateAlert(stats platform.SystemStats) {
    // Crear alerta si la CPU supera el umbral
    if(stats.CpuUsage > CPU_THRESHOLD) {
        alert := platform.Alert {
            Time: time.Now().Unix(),
            Hostname: stats.Hostname,
            Type: platform.CPUAlert,
            Value: stats.CpuUsage,
            Threshold: CPU_THRESHOLD,
            Severity: 2,
            SystemStatsID: stats.ID,
            SystemStats: stats,
        }
    }
}