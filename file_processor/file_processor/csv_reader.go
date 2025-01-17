package file_processor

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

// CSVReader implementa FileReader usando Template Method
type CSVReader struct {
	FileReaderTemplate
	Separator rune
}

// NewCSVReader inicializa CSVReader con el campo iFileReaderTemplate
func NewCSVReader(separator rune) *CSVReader {
	reader := &CSVReader{Separator: separator}
	reader.iFileReaderTemplate = reader
	return reader
}

func (r *CSVReader) parseFile(file io.Reader) ([]SiteID, error) {
	log.Println("Iniciando la lectura del archivo CSV")
	reader := csv.NewReader(file)
	reader.Comma = r.Separator

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error leyendo encabezados: %w", err)
	}

	siteIndex := -1
	idIndex := -1
	for i, header := range headers {
		if header == "site" {
			siteIndex = i
		} else if header == "id" {
			idIndex = i
		}
	}
	if siteIndex == -1 || idIndex == -1 {
		return nil, fmt.Errorf("headers 'site' o 'id' no encontrados")
	}

	var siteIDs []SiteID
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo fila: %w", err)
		}

		siteID := SiteID{Site: row[siteIndex], ID: row[idIndex]}
		siteIDs = append(siteIDs, siteID)
	}
	fmt.Println("Finalizó la lectura del archivo CSV")
	return siteIDs, nil
}
