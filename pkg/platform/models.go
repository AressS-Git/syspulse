package platform

// SystemStats define la información que los agentes instalados en los distintos equipo van a enviar al servidor principal
type SystemStats struct {
    ID uint `json:"id" gorm:"primaryKey"` // ID del equipo
    Hostname string `json:"hostname"` // Nombre del equipo
    Platform string `json:"platform"` // Plataforma del equipo (Windows o Linux)
    CpuUsage float64 `json:"cpu"` // Uso de la CPU del equipo
    RamUsage float64  `json:"ram"` // Uso de la RAM del equipo
    DiskUsage float64 `json:"disk"` // Uso del disco del equipo
    Time int64 `json:"time"` // Cuándo se generó el informe
}