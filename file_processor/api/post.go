package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"project/file_processor"
	"project/utils"
)

// SendSlicesToMediator sends the concatenated slices to the mediator service as JSON.
func SendSlicesToMediator(config utils.FileConfig, data []file_processor.SiteID) error {
	url := config.MediatorURL
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	return nil
}

// ProcessAndSendFiles selects the appropriate reader based on the format and sends the data to the mediator.
func ProcessAndSendFiles(config utils.FileConfig) error {
	reader := file_processor.NewFileReader(config)

	pattern := filepath.Join(config.Directory, "*."+config.Format)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("failed to glob files: %w", err)
	}

	var allData []file_processor.SiteID

	for _, filePath := range files {
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer f.Close()

		data, err := reader.Read(f)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		siteIDData := data
		allData = append(allData, siteIDData...)
	}

	err = SendSlicesToMediator(config, allData)
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}
