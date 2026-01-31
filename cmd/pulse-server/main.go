package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/AressS-Git/syspulse/internal/platform"
)

func main() {
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

    fmt.Println("Datos recibidos:", stats)
    writer.WriteHeader(http.StatusOK)
}