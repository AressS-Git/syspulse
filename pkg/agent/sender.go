package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AressS-Git/syspulse/pkg/platform"
)

// SendMetrics recibe los datos de collector.go y se los envía al servidor mediante HTTP POST verificando que llegaron de manera correcta
func SendMetrics(url string, stats platform.SystemStats) error {
    // statsConvertedToJson guardará las estadísticas recogidas y las convierte a JSON
    statsConvertedToJson, err := json.Marshal(stats)
    if err != nil {
        return err
    }

    // Enviar las estadístcias en formato JSON al servidor y capturar la respuesta
    serverResponse, err := http.Post(url, "application/json", bytes.NewBuffer(statsConvertedToJson))
    if err != nil {
        return err
    }

    // Cerrar el cuerpo de la respuesta por si acaso
    defer serverResponse.Body.Close()

    // Veririficar si la respuesta del servidor fue StatusOK (la establecida en el servidor cuándo todo va bien)
    if serverResponse.StatusCode != http.StatusOK {
        return fmt.Errorf("el servidor respondió con error: %v", serverResponse.Status)
    }

    return nil
}