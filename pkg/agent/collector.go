package agent

import (
	"time"
	"github.com/AressS-Git/syspulse/pkg/platform"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"
)

// GetMetrics se encargará de obtener la información necesaria de los equipos que tengan el agente instalado
func GetMetrics() (platform.SystemStats, error) {
    // Conseguir el hostname y el OS mediante el uso del paquete host y la función Info
    hostInfo, err := host.Info()
    if err != nil {
        // Si falla se devuelve la estructura vacía y el error, así con todos los campos
        return platform.SystemStats{}, err
    }

    // Conseguir el uso de la CPU
    cpuUsagePercent, err := cpu.Percent(0, false)
    if err != nil {
        return platform.SystemStats{}, err
    }

    // Conseguir el uso de la RAM
    ramUsagePercent, err := mem.VirtualMemory()
    if err != nil {
        return platform.SystemStats{}, err
    }

    // Conseguir el uso del disco
    diskUsagePercent, err := disk.Usage("/")
    if err != nil {
        return platform.SystemStats{}, err
    }

    // Obtener los procesos que más consumen gracias a la función GetProcesses
    processes, err := GetProcesses()
    if err != nil {
        return platform.SystemStats{}, err
    }

    // Obtener el tráfico de red entrante y saliente gracias a la función GetNetTraffic
    incomingNetTraffic, outboundNetTraffic, err := GetNetTraffic()
    if err != nil {
        return platform.SystemStats{}, err
    }

    // Rellenar el struct SystemStats con los datos obtenidos
    var stats platform.SystemStats
    stats.Hostname = hostInfo.Hostname
    stats.Platform = hostInfo.Platform
    stats.CpuUsage = cpuUsagePercent[0]
    stats.RamUsage = ramUsagePercent.UsedPercent
    stats.DiskUsage = diskUsagePercent.UsedPercent
    stats.Time = time.Now().Unix()
    stats.Processes = processes
    stats.IncomingNetTraffic = incomingNetTraffic
    stats.OutboundNetTraffic = outboundNetTraffic

    return stats, nil
}

// GetProcesses obtiene los cinco procesos que más recursos están consumiendo en ese momento
func GetProcesses() ([]string, error) {
    // Slice donde se guardarán los procesos que se devolverán
    var processesSlice []string
    // Obtener el slice de procesos activos
    processes, err := process.Processes()
    if err != nil {
        return nil, err
    }

    // Recorrer el slice de procesos
    for _, p := range processes {
        if len(processes) > 5 {
            break
        }
        // Obtener el nombre del proceso
        name, err := p.Name()
        if err != nil {
            return nil, err
        }
        // Añadir el proceso al slice que vamos a devolver
        processesSlice = append(processesSlice, name)
    }

    return processesSlice, nil
}

// Función que obtiene el tráfico de red total entrante y saliente
func GetNetTraffic() (uint64, uint64, error) {
    netTraffic, err := net.IOCounters(false) // False devuelve el tráfico total, true devuelve el tráfico por cada adaptador
    if err != nil {
        return 0, 0, err
    }
    // Se devuelve el tráfico entrante y saliente
    return netTraffic[0].BytesRecv, netTraffic[0].BytesSent, nil
}