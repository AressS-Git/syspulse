import { useState, useEffect } from 'react';
import './App.css';
import { GetStats } from "../wailsjs/go/main/App";

function App() {
    const [stats, setStats] = useState([]);

    useEffect(() => {
        const loadData = () => {
            GetStats().then((result) => {
                if (result) {
                    setStats(result);
                }
            });
        };

        loadData();
        const loadDataRange = setInterval(loadData, 2000);
        return () => clearInterval(loadDataRange);
    }, []);

    return (
        <div className="container">
            <h1>SYSTEM PULSE MONITOR</h1>
            
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Hostname</th>
                        <th>Plataforma</th>
                        <th>CPU</th>
                        <th>RAM</th>
                        <th>Disco</th>
                        <th>Hora</th>
                    </tr>
                </thead>
                <tbody>
                    {stats.map((item) => (
                        <tr key={item.id}>
                            <td>{item.id}</td>
                            <td>{item.hostname}</td>
                            <td>{item.platform}</td>
                            
                            {/* CAMBIO 1: Usamos .toFixed(2) para dejar solo 2 decimales y añadimos el % */}
                            <td>{item.cpu.toFixed(2)}%</td>
                            <td>{item.ram.toFixed(2)}%</td>
                            <td>{item.disk.toFixed(2)}%</td>

                            {/* CAMBIO 2: Convertimos la fecha. 
                                Multiplicamos por 1000 porque Go usa segundos y JS milisegundos */}
                            <td>{new Date(item.time * 1000).toLocaleTimeString()}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default App;