package file_processor

import (
	"io"
	"log"
	"project/utils"
)

type SiteID struct {
	Site string `json:"site"`
	ID   string `json:"id"`
}

type IFileReaderTemplate interface {
	parseFile(file io.Reader) ([]SiteID, error)
	Read(file io.Reader) ([]SiteID, error)
}

// FileReaderTemplate define la estructura del algoritmo de lectura
type FileReaderTemplate struct {
	iFileReaderTemplate IFileReaderTemplate
}

// NewFileReader inicializa el lector de archivos basado en la configuración
func NewFileReader(config utils.FileConfig) IFileReaderTemplate {
	switch config.Format {
	case "csv":
		return NewCSVReader(rune(config.Separator[0]))
	case "jsonl":
		return NewJSONLinesReader()
	default:
		log.Fatalf("Formato de archivo no soportado: %s", config.Format)
	}
	return nil
}

// Read implementa el método plantilla
func (t *FileReaderTemplate) Read(file io.Reader) ([]SiteID, error) {
	records, err := t.iFileReaderTemplate.parseFile(file)
	if err != nil {
		return nil, err
	}
	log.Println("Finalizó la lectura del archivo")
	return records, nil
}

// parseFile es un método a ser implementado por las subclases
func (t *FileReaderTemplate) parseFile(file io.Reader) ([]SiteID, error) {
	return nil, nil
}
