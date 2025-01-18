package tests

import (
	"project/api"
	"project/file_processor"
	"project/utils"
	"testing"
)

func TestSendSlicesToMediator(t *testing.T) {
	config := utils.FileConfig{
		MediatorURL: "http://localhost:8080/mediator",
	}
	data := []file_processor.SiteID{
		{Site: "MLA", ID: "839084261"},
		{Site: "MLA", ID: "813840673"},
	}

	err := api.SendSlicesToMediator(config, data)
	if err != nil {
		t.Errorf("SendSlicesToMediator failed: %v", err)
	}
}

func TestProcessAndSendFiles(t *testing.T) {
	config := utils.FileConfig{
		Directory:   "/path/to/files",
		Format:      "jsonl",
		Separator:   ",",
		Encoding:    "utf-8",
		BearerToken: "token123",
		MediatorURL: "http://localhost:8080/mediator",
	}

	err := api.ProcessAndSendFiles(config)
	if err != nil {
		t.Errorf("ProcessAndSendFiles failed: %v", err)
	}
}
