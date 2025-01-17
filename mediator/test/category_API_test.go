package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	api "project/apis"
	"testing"
)

func TestCategoryAPI_FetchData(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/categories/MLA10512/attributes" {
			t.Fatalf("Expected to request '/categories/MLA10512/attributes', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `[{"code":200, "body":{"name":"CategoryName"}}]`)
	}))
	defer server.Close()

	api := api.CategoryAPI{
		BaseURL:     server.URL,
		BearerToken: "testtoken",
	}

	result, err := api.FetchData("MLA10512")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := "CategoryName, "
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
