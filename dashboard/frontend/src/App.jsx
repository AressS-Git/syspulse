import { useState, useEffect } from 'react';
import './App.css';
// Importar las funiones de app.go para obtener los datos desde Wails
import { GetStats, GetDevices, GetAlerts } from "../wailsjs/go/main/App";

function App() {
    // Estados para almacenar los datos
    const [stats, setStats] = useState([]);
    const [devices, setDevices] = useState([]);
    const [alerts, setAlerts] = useState([]);
    
    // Estado para controlar el ID del dispositivo seleccionado
    const [selectedDeviceId, setSelectedDeviceId] = useState(null);
    // Estado para controlar qué vista estamos viendo: 'devices', 'stats' o 'alerts'
    const [view, setView] = useState('devices'); 

    useEffect(() => {
        const loadData = () => {
            if (view === 'devices') {
                GetDevices().then((result) => {
                    if (result) setDevices(result);
                });
            } else if (view === 'stats' && selectedDeviceId !== null) {
                GetStats(selectedDeviceId).then((result) => {
                    if (result) setStats(result);
                });
            } else if (view === 'alerts' && selectedDeviceId !== null) {
                // Llamamos a la nueva función de Wails
                GetAlerts(selectedDeviceId).then((result) => {
                    if (result) setAlerts(result);
                });
            }
        };

        loadData();
        const loadDataRange = setInterval(loadData, 2000);
        return () => clearInterval(loadDataRange);
    }, [selectedDeviceId, view]); // Reacciona a cambios en el ID o en la vista

    // Funciones de ayuda para cambiar de vista más limpiamente
    const showStats = (id) => {
        setSelectedDeviceId(id);
        setView('stats');
    };

    const showAlerts = (id) => {
        setSelectedDeviceId(id);
        setView('alerts');
    };

    const showDevices = () => {
        setSelectedDeviceId(null);
        setView('devices');
    };

    return (
        <div className="container">
            <h1>SYSTEM PULSE MONITOR</h1>
            {/* Tabla para mostrar los devices/dispositivos */}
            {view === 'devices' && (
                <div>
                    <h2>Equipos Detectados</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>MAC Address</th>
                                <th>Hostname</th>
                                <th>Plataforma</th>
                                <th>Acciones</th>
                            </tr>
                        </thead>
                        <tbody>
                            {devices.map((device) => (
                                <tr key={device.id}>
                                    <td>{device.id}</td>
                                    <td>{device.mac_address}</td>
                                    <td>{device.hostname}</td>
                                    <td>{device.platform}</td>
                                    <td>
                                        <button onClick={() => showStats(device.id)} style={{ marginRight: '10px' }}>
                                            Ver Estadísticas
                                        </button>
                                        {/* Nuevo botón para las alertas */}
                                        <button onClick={() => showAlerts(device.id)}>
                                            Ver Alertas
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
            {/* Tabla para mostrar las estadísticas */}
            {view === 'stats' && (
                <div>
                    <button onClick={showDevices} style={{ marginBottom: '15px' }}>
                        &larr; Volver a Equipos
                    </button>
                    <h2>Estadísticas en Vivo</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>CPU</th>
                                <th>RAM</th>
                                <th>Disco</th>
                                <th>Incoming Red Traffic</th>
                                <th>Outbound Red Traffic</th>
                                <th>Hora</th>
                            </tr>
                        </thead>
                        <tbody>
                            {stats.map((item) => (
                                <tr key={item.id}>
                                    <td>{item.id}</td>
                                    <td>{item?.cpu?.toFixed(2)}%</td>
                                    <td>{item?.ram?.toFixed(2)}%</td>
                                    <td>{item?.disk?.toFixed(2)}%</td>
                                    <td>{item.incoming_net_traffic} KB/s</td>
                                    <td>{item.outbound_net_traffic} KB/s</td>
                                    <td>{item?.time ? new Date(item.time * 1000).toLocaleTimeString() : "Cargando..."}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}

            {/* Tabla para mostrar las alertas */}
            {view === 'alerts' && (
                <div>
                    <button onClick={showDevices} style={{ marginBottom: '15px' }}>
                        &larr; Volver a Equipos
                    </button>
                    <h2>Historial de Alertas</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Tipo</th>
                                <th>Severidad</th>
                                <th>Valor Registrado</th>
                                <th>Umbral Superado</th>
                                <th>Hora</th>
                            </tr>
                        </thead>
                        <tbody>
                            {alerts.map((alert) => (
                                <tr key={alert.id}>
                                    <td>{alert.id}</td>
                                    <td>{alert.type}</td>
                                    <td>
                                        {/* Simple mapeo visual de la severidad */}
                                        {alert.severity === 1 ? 'Crítica' : alert.severity === 2 ? 'Advertencia' : 'Info'}
                                    </td>
                                    <td>{alert?.value?.toFixed(2)}</td>
                                    <td>{alert?.threshold?.toFixed(2)}</td>
                                    <td>{alert?.time ? new Date(alert.time * 1000).toLocaleTimeString() : "Cargando..."}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}
        </div>
    );
}

export default App;