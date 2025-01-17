package api_factory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ItemsAPI implementa APIClient para la API de MercadoLibre
type ItemsAPI struct {
	BaseURL     string
	BearerToken string
}

// ProductData estructura temporal para almacenar los datos necesarios
type ProductData struct {
	Price      float64 `json:"price"`
	StartTime  string  `json:"start_time"`
	CategoryID string  `json:"category_id"`
	CurrencyID string  `json:"currency_id"`
	SellerID   int64   `json:"seller_id"`
}

// FetchData obtiene datos de la API en lotes de hasta 20 productos
func (api ItemsAPI) FetchData(itemIDs []string) (map[string]interface{}, error) {
	const batchSize = 20
	var results []ProductData

	// Dividir en lotes de hasta 20 IDs
	for i := 0; i < len(itemIDs); i += batchSize {
		end := i + batchSize
		if end > len(itemIDs) {
			end = len(itemIDs)
		}
		batch := itemIDs[i:end]

		// Construir la URL para el batch
		url := fmt.Sprintf("%s/items?ids=%s", api.BaseURL, strings.Join(batch, ","))

		// Crear la solicitud GET con el Bearer Token
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error al crear solicitud HTTP: %v", err)
			return nil, fmt.Errorf("error al crear solicitud HTTP: %w", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.BearerToken))

		// Realizar la solicitud
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error al realizar solicitud a la API: %v", err)
			return nil, fmt.Errorf("error al realizar solicitud a la API: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error en respuesta de la API: %d", resp.StatusCode)
			return nil, fmt.Errorf("error en respuesta de la API: %d", resp.StatusCode)
		}

		// Decodificar la respuesta
		var response []struct {
			Code int `json:"code"`
			Body struct {
				Price      float64 `json:"price"`
				StartTime  string  `json:"start_time"`
				CategoryID string  `json:"category_id"`
				CurrencyID string  `json:"currency_id"`
				SellerID   int64   `json:"seller_id"`
			} `json:"body"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Printf("Error al decodificar respuesta: %v", err)
			return nil, fmt.Errorf("error al decodificar respuesta: %w", err)
		}

		// Extraer los campos necesarios
		for _, item := range response {
			if item.Code == http.StatusOK {
				results = append(results, ProductData{
					Price:      item.Body.Price,
					StartTime:  item.Body.StartTime,
					CategoryID: item.Body.CategoryID,
					CurrencyID: item.Body.CurrencyID,
					SellerID:   item.Body.SellerID,
				})
			}
		}

		log.Printf("Batch processed: %v", batch)
	}

	resultMap := make(map[string]interface{})
	resultMap["data"] = results
	log.Printf("FetchData completed with results: %v", results)
	return resultMap, nil
}
