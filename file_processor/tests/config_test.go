package tests

import (
	"os"
	"project/utils"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	// Setup environment variables
	os.Setenv("FILE_FORMAT", "jsonl")
	os.Setenv("FILE_SEPARATOR", ",")
	os.Setenv("FILE_ENCODING", "utf-8")
	os.Setenv("FILE_DIRECTORY", "/path/to/files")
	os.Setenv("BEARER_TOKEN", "token123")
	os.Setenv("MEDIATOR_URL", "http://localhost:8080/mediator")

	defer func() {
		// Clean up environment variables after the test
		os.Unsetenv("FILE_FORMAT")
		os.Unsetenv("FILE_SEPARATOR")
		os.Unsetenv("FILE_ENCODING")
		os.Unsetenv("FILE_DIRECTORY")
		os.Unsetenv("BEARER_TOKEN")
		os.Unsetenv("MEDIATOR_URL")
	}()

	config, err := utils.LoadConfigFromEnv()
	if err != nil {
		t.Errorf("LoadConfigFromEnv failed: %v", err)
	}

	if config.Format != "jsonl" {
		t.Errorf("Expected Format 'jsonl', got '%s'", config.Format)
	}
	if config.Separator != "," {
		t.Errorf("Expected Separator ',', got '%s'", config.Separator)
	}
	if config.Encoding != "utf-8" {
		t.Errorf("Expected Encoding 'utf-8', got '%s'", config.Encoding)
	}
	if config.Directory != "/path/to/files" {
		t.Errorf("Expected Directory '/path/to/files', got '%s'", config.Directory)
	}
	if config.BearerToken != "token123" {
		t.Errorf("Expected BearerToken 'token123', got '%s'", config.BearerToken)
	}
	if config.MediatorURL != "http://localhost:8080/mediator" {
		t.Errorf("Expected MediatorURL 'http://localhost:8080/mediator', got '%s'", config.MediatorURL)
	}
}
