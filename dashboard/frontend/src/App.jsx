import { useState, useEffect } from 'react';
import './App.css';
// Importar las funiones de app.go para obtener los datos desde Wails
import { GetStats, GetDevices, GetAlerts } from "../wailsjs/go/main/App";
// Imports para las gráficas utilizando Recharts
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';

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
            // Añadimos 'graphs' a la condición para que siga actualizando los datos en segundo plano
            } else if ((view === 'stats' || view === 'graphs') && selectedDeviceId !== null) {
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

    // Nueva función para mostrar la vista de gráficas
    const showGraphs = (id) => {
        setSelectedDeviceId(id);
        setView('graphs');
    };

    const showDevices = () => {
        setSelectedDeviceId(null);
        setView('devices');
    };

    // Preparamos los datos para las gráficas
    // Clonamos el array (con [...stats]) y le damos la vuelta para que el orden temporal vaya de izquierda a derecha.
    // También formateamos la hora para que el Eje X se lea bien.
    const graphData = [...stats].reverse().map(item => ({
        ...item,
        formattedTime: item.time ? new Date(item.time * 1000).toLocaleTimeString() : ""
    }));

    return (
        <div className="container">
            <h1>SysPulse - System Monitor</h1>
            
            {/* Tabla para mostrar los devices/dispositivos */}
            {view === 'devices' && (
                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <h2>Equipos Detectados</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Hostname</th>
                                <th>MAC Address</th>
                                <th>Plataforma</th>
                                <th>Acciones</th>
                            </tr>
                        </thead>
                        <tbody>
                            {devices.map((device) => (
                                <tr key={device.id}>
                                    <td>{device.id}</td>
                                    <td>{device.hostname}</td>
                                    <td>{device.mac_address}</td>
                                    <td>{device.platform}</td>
                                    <td>
                                        <button onClick={() => showStats(device.id)} style={{ marginRight: '10px' }}>
                                            Estadísticas
                                        </button>
                                        <button onClick={() => showAlerts(device.id)} style={{ marginRight: '10px' }}>
                                            Alertas
                                        </button>
                                        <button onClick={() => showGraphs(device.id)}>
                                            Gráficas
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
                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <button onClick={showDevices} style={{ marginBottom: '15px', alignSelf: 'flex-start', marginLeft: '2.5%' }}>
                        &larr; Volver a Equipos
                    </button>
                    <h2>Estadísticas en Vivo</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Hostname</th>
                                <th>MAC Address</th>
                                <th>CPU</th>
                                <th>RAM</th>
                                <th>Disco</th>
                                <th>Incoming Red Traffic</th>
                                <th>Outbound Red Traffic</th>
                                <th>Top Procesos (CPU)</th>
                                <th>Hora</th>
                            </tr>
                        </thead>
                        <tbody>
                            {stats.map((item) => (
                                <tr key={item.id}>
                                    <td>{item.id}</td>
                                    <td>{item.hostname}</td>
                                    <td>{item.mac_address}</td>
                                    <td>{item?.cpu?.toFixed(2)}%</td>
                                    <td>{item?.ram?.toFixed(2)}%</td>
                                    <td>{item?.disk?.toFixed(2)}%</td>
                                    <td>{(item.incoming_net_traffic / 1024).toFixed(2)} KB/s</td>
                                    <td>{(item.outbound_net_traffic / 1024).toFixed(2)} KB/s</td>
                                    <td>
                                        <div style={{ whiteSpace: 'pre-line', fontSize: '0.9em', textAlign: 'left' }}>
                                            {item.processes || "Sin datos"}
                                        </div>
                                    </td>
                                    <td>{item?.time ? new Date(item.time * 1000).toLocaleTimeString() : "Cargando..."}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}

            {/* Tabla para mostrar las alertas */}
            {view === 'alerts' && (
                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <button onClick={showDevices} style={{ marginBottom: '15px', alignSelf: 'flex-start', marginLeft: '2.5%' }}>
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
                                    <td>{alert?.value?.toFixed(2)}%</td>
                                    <td>{alert?.threshold?.toFixed(2)}%</td>
                                    <td>{alert?.time ? new Date(alert.time * 1000).toLocaleTimeString() : "Cargando..."}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}

            {/* Vista para mostrar las gráficas */}
            {view === 'graphs' && (
                <div style={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
                    <button onClick={showDevices} style={{ marginBottom: '25px', alignSelf: 'flex-start', marginLeft: '2.5%' }}>
                        &larr; Volver a Equipos
                    </button>
                    <h2 style={{ alignSelf: 'flex-start', marginLeft: '2.5%', color: '#00ADD8' }}>Monitorización Gráfica</h2>

                    {/* Gráfica de Uso de Recursos (CPU, RAM, Disco) */}
                    <div className="graph-wrapper">
                        <h3>Rendimiento del Sistema (%)</h3>
                        <ResponsiveContainer width="100%" height={300}>
                            <LineChart data={graphData}>
                                <CartesianGrid strokeDasharray="3 3" stroke="#363b45" />
                                <XAxis dataKey="formattedTime" stroke="#A0AABF" />
                                <YAxis domain={[0, 100]} stroke="#A0AABF" />
                                <Tooltip />
                                <Legend />
                                <Line type="monotone" dataKey="cpu" stroke="#ff6b6b" name="CPU Usage" strokeWidth={2} dot={false} activeDot={{ r: 8 }} />
                                <Line type="monotone" dataKey="ram" stroke="#feca57" name="RAM Usage" strokeWidth={2} dot={false} />
                                <Line type="monotone" dataKey="disk" stroke="#76AB80" name="Disk Usage" strokeWidth={2} dot={false} />
                            </LineChart>
                        </ResponsiveContainer>
                    </div>

                    {/* Gráfica de Tráfico de Red */}
                    <div className="graph-wrapper">
                        <h3>Tráfico de Red (KB/s)</h3>
                        <ResponsiveContainer width="100%" height={300}>
                            <LineChart data={graphData}>
                                <CartesianGrid strokeDasharray="3 3" stroke="#363b45" />
                                <XAxis dataKey="formattedTime" stroke="#A0AABF" />
                                <YAxis stroke="#A0AABF" />
                                <Tooltip />
                                <Legend />
                                <Line type="monotone" dataKey="incoming_net_traffic" stroke="#89E4FA" name="Incoming Traffic" strokeWidth={2} dot={false} />
                                <Line type="monotone" dataKey="outbound_net_traffic" stroke="#00ADD8" name="Outbound Traffic" strokeWidth={2} dot={false} />
                            </LineChart>
                        </ResponsiveContainer>
                    </div>
                </div>
            )}
        </div>
    );
}

export default App;