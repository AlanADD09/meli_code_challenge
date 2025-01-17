// config.go
package utils

import (
	"fmt"
	"os"
)

// FileConfig define la configuración para la lectura de archivos
type FileConfig struct {
	Format      string
	Separator   string
	Encoding    string
	Directory   string
	BearerToken string
}

// LoadConfigFromEnv carga la configuración desde el archivo .env
func LoadConfigFromEnv() (FileConfig, error) {
	envFile, err := os.Open("config-csv.env")
	if err != nil {
		return FileConfig{}, fmt.Errorf("error al abrir archivo .env: %w", err)
	}
	defer envFile.Close()

	config := FileConfig{
		Format:      os.Getenv("FILE_FORMAT"),
		Separator:   os.Getenv("FILE_SEPARATOR"),
		Encoding:    os.Getenv("FILE_ENCODING"),
		Directory:   os.Getenv("FILE_DIRECTORY"),
		BearerToken: os.Getenv("BEARER_TOKEN"),
	}

	// Validar que los campos esenciales no estén vacíos
	if config.Format == "" || config.Separator == "" || config.Encoding == "" || config.Directory == "" || config.BearerToken == "" {
		return FileConfig{}, fmt.Errorf("configuración incompleta en .env")
	}

	return config, nil
}
