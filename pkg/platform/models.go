package platform

// Tipo de Alerta
type AlertType string

const (
    CPUAlert AlertType = "CPU_HIGH_USAGE"
    RAMAlert AlertType = "RAM_HIGH_USAGE"
    DISKAlert AlertType = "DISK_HIGH_USAGE"
)

// SystemStats define la información que los agentes instalados en los distintos equipo van a enviar al servidor principal
type SystemStats struct {
    ID uint `json:"id" gorm:"primaryKey"` // ID del struct
    DeviceID uint `json:"device_id" gorm:"index"` // Clave foránea al Device
    MacAddress string `json:"mac_address"` // Dirección MAC del equipo
    Hostname string `json:"hostname"` // Nombre del equipo
    Platform string `json:"platform"` // Plataforma del equipo
    CpuUsage float64 `json:"cpu"` // Uso de la CPU del equipo
    RamUsage float64  `json:"ram"` // Uso de la RAM del equipo
    DiskUsage float64 `json:"disk"` // Uso del disco del equipo
    IncomingNetTraffic int64 `json:"incoming_net_traffic"` // Tráfico de red entrante
    OutboundNetTraffic int64 `json:"outbound_net_traffic"` // Tráfico de red saliente
    Processes string `json:"processes"` // Top 5 procesos que más están consumiendo
    Time int64 `json:"time"` // Cuándo se generó el informe
}

// Alert define la información que el servidor enviará a los usuarios para notificar cualquier irregularidad detectada en los equipos por el agente
type Alert struct {
    ID uint `json:"id" gorm:"primaryKey"` // ID de la alerta
    DeviceID uint `json:"device_id" gorm:"index"` // Clave foránea al Device
    Time int64   `json:"time"` // Cuándo ocurrió la alerta
    Hostname string `json:"hostname" gorm:"index"` // Nombre del equipo
    Type AlertType `json:"type" gorm:"index;type:text"` // Tipo de alerta
    Value float64 `json:"value"` // Valor de la alerta
    Threshold float64 `json:"threshold"` // Umbral de la alerta
    Severity uint8 `json:"severity"` // Severidad de la alerta (Info, Warning o Critical)
    SystemStatsID uint `json:"system_stats_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // ID de las estadísticas del equipo
    SystemStats SystemStats `json:"system_stats"` // Información de las estadísticas del equipo
}

// Device sirve para guardar cada dispositivo que entra en contacto con el servidor, esto permite poder agrupar las estadísticas por equipos
type Device struct {
    ID uint `json:"id" gorm:"primaryKey"`
    MacAddress string `json:"mac_address" gorm:"uniqueIndex"` // Dirección MAC única para identificar al equipo
    Hostname string `json:"hostname"` // Único para evitar duplicados
    Platform string `json:"platform"` // Plataforma del equipo
    
    // Relaciones 1:N con las estadísticas y las alertas
    SystemStats []SystemStats `json:"system_stats" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    Alerts      []Alert       `json:"alerts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}