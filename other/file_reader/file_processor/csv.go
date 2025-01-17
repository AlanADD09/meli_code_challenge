// reader_factory.go
package file_processor

import (
	"encoding/csv"
	"fmt"
	"io"
)

// CSVReader implementa FileReader para archivos CSV
type CSVReader struct {
	Separator rune
}

func (r CSVReader) Read(file io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(file)
	reader.Comma = r.Separator

	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error leyendo encabezados: %w", err)
	}

	var records []map[string]string
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error leyendo fila: %w", err)
		}

		record := map[string]string{}
		for i, value := range row {
			record[headers[i]] = value
		}
		records = append(records, record)
	}
	return records, nil
}
