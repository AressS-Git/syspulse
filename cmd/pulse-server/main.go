package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/AressS-Git/syspulse/pkg/platform"
	"github.com/AressS-Git/syspulse/pkg/server"
)


func main() {
    // Se inicia la conexión con la BD, si da error, el servidor se detendrá graicas a los panics de la función
    server.InitDB()

    fmt.Println("Conexión a la BD establecida correctamente y tablas creadas correctamente")

    // Las peticiones http entrantes que accedan a dicha ruta se manejarán con el handler
    http.HandleFunc("/api/stats", httpHandler)

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

    // Guardar los datos de stats en la BD, usamos la variable global del otro paquete
    result := server.DB.Create(&stats)
    if result.Error != nil {
        http.Error(writer, "Error al guardar los datos en la BD", http.StatusInternalServerError)
    }

    fmt.Println("Datos recibidos y guardados en la BD:", stats.ID, stats)
    writer.WriteHeader(http.StatusOK)
}