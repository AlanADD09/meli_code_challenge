package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	api "project/apis"
	"testing"
)

func TestCurrencyAPI_FetchData(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/currencies/ARS" {
			t.Fatalf("Expected to request '/currencies/ARS', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"code":200, "body":{"description":"Argentine Peso"}}`)
	}))
	defer server.Close()

	api := api.CurrencyAPI{
		BaseURL:     server.URL,
		BearerToken: "testtoken",
	}

	result, err := api.FetchData("ARS")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "Argentine Peso"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
