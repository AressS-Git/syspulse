package server

import (
    "fmt"
    "encoding/json"
    "os"
    "net/smtp"
    "strings"
    "time"
    "github.com/AressS-Git/syspulse/pkg/platform"
)

// EmailNotifier es la estructura de configuración del servidor de correo
type EmailNotifier struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
    To       []string
}

// LoadNotifierConfig obtiene la configuración del servidor de correo de un archivo JSON
func LoadNotifierConfig(filename string) (*EmailNotifier, error) {
    // Leer el archivo de configuración
    notifierConfigFile, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("Error al intentar leer el archivo de configuración: %v", err)
    }

    // Decodificar el archivo JSON en una estructura tipo EmailNotifier
    var notifierConfig EmailNotifier
    if err := json.Unmarshal(notifierConfigFile, &notifierConfig); err != nil {
        return nil, fmt.Errorf("Error al decodificar el archivo JSON de configuración: %v", err)
    }
    return &notifierConfig, nil
}

// Notifiy envía el correo a todas las direcciones destinatarias establecidas en la configuración del servidor de correo
func (notifier *EmailNotifier) Notify(message string, body string) error {
    // Crear la dirección del servidor (host + puerto)
    serverAddress := fmt.Sprintf("%v:%v", notifier.Host, notifier.Port)
    
    // Configurar la autorización utilizando la librería smtp
    auth := smtp.PlainAuth(
        "",
        notifier.Username,
        notifier.Password,
        notifier.Host,
    )

    // Construir el mensaje(cabecera + contenido)
    // En los correos se usan retornos diferentes r y n
    // Cabecera
    var header string
    header += fmt.Sprintf("From: %v\r\n", notifier.From)
    header += fmt.Sprintf("To: %v\r\n", strings.Join(notifier.To, ", "))
    header += fmt.Sprintf("Subject: %v\r\n", message)

    // Añadir cabeceras HTML para que el mensaje tenga formato HTML y así sea más atractivo
    header += "MIME-version: 1.0;\r\n"
    header += "Content-Type: text/html; charset=\"UTF-8\";\r\n"

    // Separar cabecera de cuerpo (línea en blanco)
    header += "\r\n"

    // Cuerpo
    header += fmt.Sprintf("%v\r\n", body)

    // Enviar el correo
    if err := smtp.SendMail(
        serverAddress,
        auth,
        notifier.From,
        notifier.To,
        []byte(header), // El mensaje se envía en Bytes cómo el TCP en Java
    ); err != nil {
        return fmt.Errorf("Error al enviar el correo: %v", err)
    }
    return nil
}

// StartPeriodicReport envía un reporte con las últimas alertas cada cierto intervalo de tiempo
func StartPeriodicReport(notifier *EmailNotifier, interval time.Duration) {
    if notifier == nil {
        fmt.Println("El servicio de correo no está configurado. Se cancela el reporte periódico.")
        return
    }

    for {
        var alertas []platform.Alert
        
        // Extraer las últimas 20 alertas
        if err := DB.Order("time desc").Limit(20).Find(&alertas).Error; err != nil {
            fmt.Println("Error al extraer las alertas periódicas:", err)
            // IMPORTANTE: Dormir también si hay error para evitar un bucle infinito que sature la CPU
            time.Sleep(interval)
            continue
        }

        // Inicio del cuerpo HTML
        body := `
        <!DOCTYPE html>
        <html>
        <head>
            <style>
                table { width: 100%; border-collapse: collapse; font-family: Arial, sans-serif; }
                th, td { border: 1px solid #dddddd; text-align: left; padding: 8px; font-size: 14px; }
                th { background-color: #2c3e50; color: white; }
                tr:nth-child(even) { background-color: #f2f2f2; }
                .high-severity { color: #e74c3c; font-weight: bold; }
            </style>
        </head>
        <body>
            <h2>Reporte Automático de Alertas SysPulse</h2>
        `

        if len(alertas) == 0 {
            body += "<p>Actualmente no hay ninguna alerta registrada en el sistema.</p>"
        } else {
            // Crear los encabezados de la tabla
            body += `
            <table>
              <tr>
                <th>Host</th>
                <th>Tipo</th>
                <th>Severidad</th>
                <th>Valor</th>
                <th>Umbral</th>
              </tr>
            `
            // Llenar la tabla con las alertas
            for _, alerta := range alertas {
                
                // Pasar el campo serverity uint8 a texto y asignar un estilo a cada tipo de severity
                var sevLabel string
                var sevStyle string

                // Cada severity 
                switch alerta.Severity {
                case 3:
                    sevLabel = "Info"
                    sevStyle = ` style="color: #27ae60;"` // Verde
                case 2:
                    sevLabel = "Warning"
                    sevStyle = ` style="color: #f39c12; font-weight: bold;"` // Naranja
                case 1:
                    sevLabel = "Critical"
                    sevStyle = ` class="high-severity"` // Rojo
                default:
                    sevLabel = fmt.Sprintf("Desconocida (%d)", alerta.Severity)
                }

                // Añadir los valores (incluido sevLabel y sevStyle)
                body += fmt.Sprintf(`
                <tr>
                    <td>%v</td>
                    <td>%v</td>
                    <td%v>%v</td>
                    <td>%.2f</td>
                    <td>%.2f</td>
                </tr>`, 
                alerta.Hostname, alerta.Type, sevStyle, sevLabel, alerta.Value, alerta.Threshold)
            }
            body += "</table>" // Cerrar la tabla
        }
        
        body += "</body></html>" // Cerrar el cuerpo HTML

        err := notifier.Notify("Reporte SysPulse: Alertas Periódicas", body)
        if err != nil {
            fmt.Println("Error al enviar el correo periódico:", err)
        } else {
            fmt.Println("Correo periódico enviado con éxito a", notifier.To)
        }

        // Pausar la ejecución durante el intervalo establecido
        // Al final de la función para que se ejecute al iniciar el servidor y no espere 6 horas hasta hacerlo, así es más rentable para las pruebas
        time.Sleep(interval)
    }
}