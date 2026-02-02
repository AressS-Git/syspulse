import { useState } from 'react';
import './App.css';

function App() {
    // Aquí guardaremos los datos que vengan de Go
    const [stats, setStats] = useState([]);

    return (
        <div className="container">
            <h1>SYSTEM PULSE MONITOR</h1>
            
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Hostname</th>
                        <th>Plataforma</th>
                        <th>CPU (%)</th>
                        <th>RAM (%)</th>
                        <th>Disco (%)</th>
                    </tr>
                </thead>
                <tbody>
                    {/* Aquí pintaremos los datos dinámicamente. 
                        De momento ponemos una fila de ejemplo fake */}
                    <tr>
                        <td>1</td>
                        <td>Server-Alpha</td>
                        <td>linux</td>
                        <td>12.5%</td>
                        <td>45.2%</td>
                        <td>80.1%</td>
                    </tr>
                    <tr>
                        <td>2</td>
                        <td>Laptop-Admin</td>
                        <td>darwin</td>
                        <td>5.0%</td>
                        <td>60.0%</td>
                        <td>20.0%</td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
}

export default App;