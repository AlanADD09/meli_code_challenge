package api_factory

import "fmt"

// APIClient define la interfaz para interactuar con una API externa
type APIClient interface {
	FetchData(itemIDs []string) (map[string]interface{}, error)
}

// GetAPIClient retorna el cliente adecuado según el tipo de API configurada
func GetAPIClient(apiType string, baseURL string, bearerToken string) (APIClient, error) {
	switch apiType {
	case "mercadolibre":
		return ItemsAPI{BaseURL: baseURL, BearerToken: bearerToken}, nil
	default:
		return nil, fmt.Errorf("tipo de API no soportado: %s", apiType)
	}
}
