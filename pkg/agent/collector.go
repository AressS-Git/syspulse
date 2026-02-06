package agent

import (
	"time"
	"github.com/AressS-Git/syspulse/pkg/platform"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// GetMetrics se encargará de obtener la información necesaria de los equipos que tengan el agente instalado
func GetMetrics() (platform.SystemStats, error) {
    // 1. Conseguir el hostname y el OS mediante el uso del paquete host y la función Info
    hostInfo, err := host.Info()
    if err != nil {
        // Si falla se devuelve la estructura vacía y el error, así con todos los campos
        return platform.SystemStats{}, err
    }

    // 2. Conseguir el uso de la CPU
    cpuUsagePercent, err := cpu.Percent(0, false)
    if err != nil {
        return platform.SystemStats{}, err
    }

    // 3. Conseguir el uso de la RAM
    ramUsagePercent, err := mem.VirtualMemory()
    if err != nil {
        return platform.SystemStats{}, err
    }

    // 4. Conseguir el uso del disco
    diskUsagePercent, err := disk.Usage("/")
    if err != nil {
        return platform.SystemStats{}, err
    }

    // 5. Rellenar el struct SystemStats con los datos obtenidos
    var stats platform.SystemStats
    stats.Hostname = hostInfo.Hostname
    stats.Platform = hostInfo.Platform
    stats.CpuUsage = cpuUsagePercent[0]
    stats.RamUsage = ramUsagePercent.UsedPercent
    stats.DiskUsage = diskUsagePercent.UsedPercent
    stats.Time = time.Now().Unix()

    return stats, nil
}