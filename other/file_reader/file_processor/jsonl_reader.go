// reader_factory.go
package file_processor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// JSONLinesReader implementa FileReader para archivos JSONLines
type JSONLinesReader struct{}

func (r JSONLinesReader) Read(file io.Reader) ([]map[string]string, error) {
	scanner := bufio.NewScanner(file)
	var records []map[string]string

	for scanner.Scan() {
		line := scanner.Text()
		var record map[string]string
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, fmt.Errorf("error leyendo línea JSON: %w", err)
		}
		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error escaneando archivo JSONLines: %w", err)
	}

	return records, nil
}
