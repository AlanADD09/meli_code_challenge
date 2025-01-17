// config.go
package utils

import (
	"fmt"
	"os"
)

// FileConfig define la configuración para la lectura de archivos
type FileConfig struct {
	MeliURL     string
	BearerToken string
	// MediatorURL string
}

// LoadConfigFromEnv carga la configuración desde el archivo .env
func LoadConfigFromEnv() (FileConfig, error) {
	envFile, err := os.Open("config.env")
	if err != nil {
		return FileConfig{}, fmt.Errorf("error al abrir archivo .env: %w", err)
	}
	defer envFile.Close()

	config := FileConfig{
		MeliURL:     os.Getenv("MELI_URL"),
		BearerToken: os.Getenv("BEARER_TOKEN"),
	}

	// Validar que los campos esenciales no estén vacíos
	if config.MeliURL == "" || config.BearerToken == "" {
		return FileConfig{}, fmt.Errorf("configuración incompleta en .env")
	}

	return config, nil
}
