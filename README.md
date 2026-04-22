# SysPulse: Remote System Monitoring Tool

**SysPulse** es una herramienta de monitorización de sistemas ligera y eficiente diseñada para administradores que buscan sencillez y potencia. El sistema utiliza agentes ligeros en Go para recolectar métricas en tiempo real de servidores remotos y visualizarlas en un dashboard centralizado nativo para macOS.

---

## 🇪🇸 Contenido en Español

### Descripción General
SysPulse resuelve el problema del acceso remoto a las métricas de rendimiento sin necesidad de complejas configuraciones de red o pesados stacks de software. Proporciona una visión inmediata del "pulso" del servidor mediante una arquitectura Agente-Servidor altamente concurrente.

### Características Principales
* **Monitorización Centralizada**: Control de múltiples nodos (Linux/macOS) desde una única interfaz.
* **Visualización en Tiempo Real**: Gráficos dinámicos y reactivos utilizando Recharts.
* **Alertas Inteligentes**: Notificaciones automáticas vía correo electrónico (SMTP) basadas en umbrales de CPU/RAM.
* **Persistencia Local**: Almacenamiento histórico mediante SQLite y el ORM GORM.
* **Bajo Consumo**: Agentes diseñados para tener un impacto mínimo en el rendimiento del sistema monitorizado.

### Tecnologías Utilizadas
* **Backend & Agentes**: Go (Golang).
* **Frontend**: React.js y Recharts.
* **Framework Desktop**: Wails (puente nativo entre Go y el motor web).
* **Base de Datos**: SQLite.

### Instalación y Ejecución

#### 1. Preparación
Asegúrate de tener instalado Go, Node.js y la CLI de Wails (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`).

#### 2. Compilación del proyecto
Para generar los ejecutables finales de forma automática, ejecuta el script de construcción:
```bash
./build.sh
```

#### 3. Lanzar el Servidor (macOS)
Ejecuta el binario generado para abrir el Dashboard central:
```bash
./sys-server
```

#### 4. Conectar un Agente (Linux / macOS)
En la máquina remota que quieras monitorizar, lanza el agente indicando la URL del servidor (normalmente puerto 8080):
```bash
./pulse-agent -url http://[IP_DE_TU_MAC]:8080/api/stats
```

---

## 🇺🇸 English Content

### Project Overview
SysPulse is a lightweight system monitoring tool built for administrators who prioritize efficiency. It consists of Go-based agents that gather real-time hardware metrics from remote servers and display them on a modern, centralized macOS dashboard.

### Key Features
* **Centralized Monitoring**: Manage multiple remote nodes from a single desktop interface.
* **Real-time Visualization**: Reactive charts providing immediate feedback on system health.
* **Smart Alerting**: Automatic email notifications (SMTP) triggered by customizable performance thresholds.
* **Historical Data**: Persistent storage using a zero-config SQLite database.
* **Resilient Architecture**: Agents feature automatic reconnection logic and low resource overhead.

### Tech Stack
* **Core**: Go (Golang) for high-performance backend and agents.
* **UI**: React.js & Recharts for a component-based, modern dashboard.
* **Bridge**: Wails Framework (connecting Go logic with the Web frontend).
* **Storage**: SQLite with GORM.

### Installation & Deployment

#### 1. Requirements
Ensure you have Go, Node.js, and the Wails CLI installed on your development machine.

#### 2. Building the Project
Run the provided build script to compile binaries for both the server and the agents:
```bash
./build.sh
```

#### 3. Starting the Server (macOS)
Run the server binary to launch the GUI dashboard:
```bash
./sys-server
```

#### 4. Starting the Agent (Linux / macOS)
Deploy the agent to remote machines and connect it to the server's IP address:
```bash
./pulse-agent -url http://[YOUR_MAC_IP]:8080/api/stats
```

---

### 📂 Estructura del Proyecto / Project Structure
* `cmd/`: Entry points for Agent and Server.
* `pkg/`: Core logic (collector, database, alerts, handlers).
* `dashboard/`: UI source code (Wails/React/Vite).
* `platform/`: Shared data models (SystemStats, Device).

### 🛠 Trabajos Futuros / Future Work
* Windows Agent support.
* Advanced security via JWT tokens for Agent-Server communication.
* Detailed Disk and Network I/O charts.

---
**Desarrollado por / Developed by**: Sergio Gómez Santos  
**Proyecto Final de Ciclo**: DAM (Desarrollo de Aplicaciones Multiplataforma)
