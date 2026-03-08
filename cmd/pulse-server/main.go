package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/AressS-Git/syspulse/pkg/platform"
    "github.com/AressS-Git/syspulse/pkg/server"
)

// Variable para acceder al notificador desde el handler
var emailNotifier *server.EmailNotifier

func main() {
    // Cargar la configuración de las alertas desde el archivo alertsparameters.json
    server.CreateThresholds()

    // Se inicia la conexión con la BD, si da error, el servidor se detendrá gracias a los panics de la función
    server.InitDB()

    fmt.Println("Conexión a la BD establecida correctamente y tablas creadas correctamente")

    // Iniciar el notficador
    var errNotifier error
    // Usar la función LoadNotifierConfig para cargar la configuración del servidor de correo desde el JSON
    emailNotifier, errNotifier = server.LoadNotifierConfig("pkg/server/notifierparameters.json")
    if errNotifier != nil {
        fmt.Println("Error al cargar la configuración del correo:", errNotifier)
    } else {
        fmt.Println("Notificador de correo configurado correctamente")
        
        // Iniciar el envío automático de reportes en segundo plano cada 6 horas
        go server.StartPeriodicReport(emailNotifier, 6*time.Hour)
    }

    // Las peticiones http entrantes que accedan a dicha ruta se manejarán con el handler
    http.HandleFunc("/api/stats", httpHandler)

    // Ruta para disparar el envío del correo manualmente
    http.HandleFunc("/api/report", reportHandler)

    // Imprimir antes de que el servidor bloquee la ejecución
    fmt.Println("Servidor escuchando en http://localhost:8080/api/stats...")

    // Abrir el puerto y escuchar peticiones
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error al iniciar el servidor:", err)
    }
}

// httpHandler maneja las peticiones http que lleguen por el puerto (writer escribe y request representa la petición)
func httpHandler(writer http.ResponseWriter, request *http.Request) {
    // Cerrar el body de las request
    defer request.Body.Close()

    if request.Method != http.MethodPost {
        http.Error(writer, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    // stats guardará los datos que saquemos de la request
    var stats platform.SystemStats

    // Guardar los datos de la request en stats
    err := json.NewDecoder(request.Body).Decode(&stats)
    if err != nil {
        http.Error(writer, "JSON no válido", http.StatusBadRequest)
        return
    }

    // Buscar el dispositivo por su MAC o crearlo si es la primera vez que se conecta
    var device platform.Device
    errDevice := server.DB.Where(platform.Device{MacAddress: stats.MacAddress}).FirstOrCreate(&device, platform.Device{
        MacAddress: stats.MacAddress,
        Hostname:   stats.Hostname,
        Platform:   stats.Platform,
    }).Error

    if errDevice != nil {
        http.Error(writer, "Error al gestionar el dispositivo en la BD", http.StatusInternalServerError)
        return
    }

    // Asignar el ID del dispositivo a las estadísticas generadas
    stats.DeviceID = device.ID

    // Guardar los datos de stats en la BD, usamos la variable global del otro paquete
    result := server.DB.Create(&stats)
    if result.Error != nil {
        http.Error(writer, "Error al guardar los datos en la BD", http.StatusInternalServerError)
        return
    }
    
    fmt.Println("Datos guardados en la BD, creando alerta en segundo plano...")
    
    // Goroutine para que la alerta se vaya creando de fondo y se pueda ir devolviendo la respuesta sin tener que esperar a que la BD responda al crear la alerta
    // La función CreateAlert se encargará de crear la alerta en base a los umbrales establecidos en alerts.go
    go server.CreateAlert(stats)
    
    writer.WriteHeader(http.StatusOK)
}

// reportHandler extrae alertas de la BD y las envía por correo
func reportHandler(writer http.ResponseWriter, request *http.Request) {
    // Verificar que el notificador se cargó correctamente al inicio
    if emailNotifier == nil {
        http.Error(writer, "El servicio de correo no está configurado", http.StatusInternalServerError)
        return
    }

    var alertas []platform.Alert
    
    // Extraer las últimas 20 alertas de la base de datos ordenadas por fecha descendente
    if err := server.DB.Order("time desc").Limit(20).Find(&alertas).Error; err != nil {
        http.Error(writer, "Error al extraer las alertas de la base de datos", http.StatusInternalServerError)
        return
    }

    // Construir el cuerpo del correo
    body := "Reporte de las últimas alertas del sistema (Max 20):\n\n"
    if len(alertas) == 0 {
        body += "Actualmente no hay ninguna alerta registrada en el sistema.\n"
    } else {
        for _, alerta := range alertas {
            // Se formatea la información de cada alerta en una línea
            body += fmt.Sprintf("- Host: %v | Tipo: %v | Sev: %v | Valor: %.2v | Umbral: %.2v\n",
                alerta.Hostname, alerta.Type, alerta.Severity, alerta.Value, alerta.Threshold)
        }
    }

    // Enviar el correo en una goroutine mediante una función anónima para no bloquear la respuesta HTTP
    go func() {
        err := emailNotifier.Notify("Reporte SysPulse: Últimas Alertas", body)
        if err != nil {
            fmt.Println("Error al enviar el correo de reporte:", err)
        } else {
            fmt.Println("Correo de reporte de alertas enviado con éxito a", emailNotifier.To)
        }
    }()

    writer.WriteHeader(http.StatusOK)
    if _, err := writer.Write([]byte("Solicitud recibida. Procesando el envío del correo en segundo plano...")); err != nil {
        fmt.Println("Error al enviar la respuesta al cliente:", err)
    }
}