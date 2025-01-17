package file_processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// JSONLinesReader implementa FileReader usando Template Method
type JSONLinesReader struct {
	FileReaderTemplate
}

// NewJSONLinesReader initializes JSONLinesReader with the required interface
func NewJSONLinesReader() *JSONLinesReader {
	reader := &JSONLinesReader{}
	reader.iFileReaderTemplate = reader
	return reader
}

func (r *JSONLinesReader) parseFile(file io.Reader) ([]SiteID, error) {
	fmt.Println("Iniciando la lectura del archivo JSONLines")
	scanner := bufio.NewScanner(file)
	var siteIDs []SiteID

	for scanner.Scan() {
		line := scanner.Text()
		var record SiteID
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, fmt.Errorf("error leyendo línea JSON: %w", err)
		}

		siteIDs = append(siteIDs, record)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error escaneando archivo JSONLines: %w", err)
	}

	fmt.Println("Finalizó la lectura del archivo JSONLines")

	return siteIDs, nil
}
