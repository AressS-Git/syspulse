package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "SysPulse Server OK")
    })

    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}