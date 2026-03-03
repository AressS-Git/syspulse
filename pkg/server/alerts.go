package server

import (
    "github.com/AressS-Git/syspulse/pkg/platform"
    "fmt"
    "time"
)

// Threshold sirve cómo molde de cada uno de los umbrales que se usarán y del impacto que tiene superar cada uno de ellos
type Threshold struct {
    Value float64
    Severity uint8
}

// Umbrales ordenados de más grave (1) a menos grave (3)
var cpuThresholds = []Threshold{
    {95.0, 1},
    {90.0, 2},
    {80.0, 3},
}

var ramThresholds = []Threshold{
    {100.0, 1},
    {95.0, 2},
    {85.0, 3},
}

var diskThresholds = []Threshold{
    {100.0, 1},
    {85.0, 2},
    {70.0, 3},
}

// CreateAlert llamará a la función ChechkStats pasándola los distintos parámetos que debe revisar y los umbrales que indican si hay que crear una laerta o no es necesario
func CreateAlert(stats platform.SystemStats) {
    CheckStats(stats, stats.CpuUsage, cpuThresholds, platform.CPUAlert)
    CheckStats(stats, stats.RamUsage, ramThresholds, platform.RAMAlert)
    CheckStats(stats, stats.DiskUsage, diskThresholds, platform.DISKAlert)
}

// CheckStats comprobará si las estadísticas superan los umbrales o no, si es así creará una alerta en la BD
func CheckStats(stats platform.SystemStats, value float64, thresholds []Threshold, alertType platform.AlertType) {
    for _, threshold := range thresholds {
        // Si el valor recogido por el agente supera el umbral se creará una alerta del tipo concreto en cada caso
        if value > threshold.Value {
            alert := platform.Alert{
                Time:          time.Now().Unix(),
                Hostname:      stats.Hostname,
                Type:          alertType,
                Value:         value,
                Threshold:     threshold.Value,
                Severity:      threshold.Severity,
                SystemStatsID: stats.ID,
                SystemStats:   stats,
            }
    
            // Se añade la alerta a la BD
            if err := DB.Create(&alert).Error; err != nil {
                fmt.Println("Error al guardar la alerta:", err)
            } else {
                fmt.Println("Alerta creada:", alert)
            }
        }
    }
}