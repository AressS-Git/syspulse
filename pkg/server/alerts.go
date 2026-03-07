package server

import (
    "github.com/AressS-Git/syspulse/pkg/platform"
    "fmt"
    "time"
    "encoding/json"
    "os"
)

// Threshold sirve cómo molde de cada uno de los umbrales que se usarán y del impacto que tiene superar cada uno de ellos
type Threshold struct {
    Value    float64 `json:"value"`
    Severity uint8   `json:"severity"`
}

type ConfigAlerts struct {
    CPUThresholds  []Threshold `json:"cpu_thresholds"`
    RAMThresholds  []Threshold `json:"ram_thresholds"`
    DiskThresholds []Threshold `json:"disk_thresholds"`
}

// Umbrales ordenados de más grave (1) a menos grave (3)
var cpuThresholds = []Threshold{
    {100.0, 1},
    {75.0, 2},
    {50.0, 3},
}

var ramThresholds = []Threshold{
    {100.0, 1},
    {80.0, 2},
    {60.0, 3},
}

var diskThresholds = []Threshold{
    {100.0, 1},
    {75.0, 2},
    {50.0, 3},
}

// CreatThresholds que extrae los umbrales de las alertas alojados en el archivo alertsparameters.json
// Si falla utiliza los valores por defecto
func CreateThresholds() {
    // Archivo a decodificar
    const filename = "pkg/server/alertsparameters.json"

    // Leer el archivo
    fileInfo, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error al leer el JSON:", err)
        return
    }

    // Decodificar el JSON y guardar el resultado en una variable ConfigAlerts
    var config ConfigAlerts
    err = json.Unmarshal(fileInfo, &config)
    if err != nil {
        fmt.Println("Error al decodificar el JSON:", err)
        return
    }

    // Si la variable config ha recogido los datos del JSON estos se utilizarán para establecer los umbrales
    if len(config.CPUThresholds) > 0 {
        cpuThresholds = config.CPUThresholds
    }
    if len(config.RAMThresholds) > 0 {
        ramThresholds = config.RAMThresholds
    }
    if len(config.DiskThresholds) > 0 {
        diskThresholds = config.DiskThresholds
    }

    fmt.Println("Umbrales cargados correctamente:", config)
}

// CreateAlert llamará a la función ChechkStats pasándola los distintos parámetos que debe revisar y los umbrales que indican si hay que crear una laerta o no es necesario
func CreateAlert(stats platform.SystemStats) {
    // Llamar a la función CheckStats para que compruebe cada estadística para saber si necesita una alerta
    CheckStats(stats, stats.CpuUsage, cpuThresholds, platform.CPUAlert)
    CheckStats(stats, stats.RamUsage, ramThresholds, platform.RAMAlert)
    CheckStats(stats, stats.DiskUsage, diskThresholds, platform.DISKAlert)
}

// CheckStats comprobará si las estadísticas superan los umbrales o no, si es así creará una alerta en la BD
func CheckStats(stats platform.SystemStats, value float64, thresholds []Threshold, alertType platform.AlertType) {
    for _, threshold := range thresholds {
        // Si el valor recogido por el agente supera el umbral se creará una alerta del tipo concreto en cada caso
        if value >= threshold.Value {
            alert := platform.Alert{
                DeviceID:      stats.DeviceID,
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
            break
        }
    }
}