// reader_factory.go
package file_processor

import (
	"fmt"
	"io"
	"path/filepath"
	"project/utils"
)

// FileReader define la interfaz para leer archivos
type FileReader interface {
	Read(file io.Reader) ([]map[string]string, error)
}

// GetReader retorna el lector adecuado según el formato configurado
func GetReader(filePath string, config utils.FileConfig) (FileReader, error) {
	// Detectar la extensión del archivo
	ext := filepath.Ext(filePath)

	// Seleccionar el lector basado en la extensión
	switch ext {
	case ".csv":
		separator := ','
		if len(config.Separator) > 0 {
			separator = rune(config.Separator[0])
		}
		return CSVReader{Separator: separator}, nil
	case ".jsonl", ".jsonlines":
		return JSONLinesReader{}, nil
	default:
		// Si no se reconoce la extensión, usar el formato configurado como respaldo
		switch config.Format {
		case "csv":
			separator := ','
			if len(config.Separator) > 0 {
				separator = rune(config.Separator[0])
			}
			return CSVReader{Separator: separator}, nil
		case "jsonlines":
			return JSONLinesReader{}, nil
		default:
			return nil, fmt.Errorf("formato no soportado: extensión %s o configuración %s", ext, config.Format)
		}
	}
}
