package api

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	// "project/mediator"

	"strings"
)

type ItemsAPI struct {
	// mediator     mediator.Mediator
	BaseURL     string
	BearerToken string
}

func (api *ItemsAPI) FetchData(itemID string) (string, error) {
	return "", nil
}

func (api *ItemsAPI) FetchNumericData(input int64) (string, error) {
	return "", nil
}

func (api *ItemsAPI) FetchMultiData(itemIDs []SiteID) ([]ResponseData, error) {
	const batchSize = 20
	var results []ResponseData
	fmt.Println("Fetching data for items")
	// Dividir en lotes de hasta 20 IDs
	for i := 0; i < len(itemIDs); i += batchSize {
		end := i + batchSize
		if end > len(itemIDs) {
			end = len(itemIDs)
		}
		batch := itemIDs[i:end]

		var ids []string
		for _, id := range batch {
			ids = append(ids, fmt.Sprintf("%s%s", id.Site, id.ID))
		}

		// Construir la URL para el batch
		url := fmt.Sprintf("%s/items?ids=%s", api.BaseURL, strings.Join(ids, ","))

		fmt.Println("URL: ", url)

		// Crear la solicitud GET con el Bearer Token
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Printf("Error al crear solicitud HTTP: %v", err)
			return nil, fmt.Errorf("error creating HTTP request: %v", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.BearerToken))

		// Realizar la solicitud
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error al realizar solicitud a la API: %v", err)
			return nil, fmt.Errorf("error making API request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error en respuesta de la API: %d", resp.StatusCode)
			return nil, fmt.Errorf("API response error: %d", resp.StatusCode)
		}

		// Decodificar la respuesta
		var response []struct {
			Code int `json:"code"`
			Body struct {
				SiteID     string  `json:"site_id"`
				ID         string  `json:"id"`
				Price      float64 `json:"price"`
				StartTime  string  `json:"date_created"`
				CategoryID string  `json:"category_id"`
				CurrencyID string  `json:"currency_id"`
				SellerID   int64   `json:"seller_id"`
			} `json:"body"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Printf("Error al decodificar respuesta: %v", err)
			return nil, fmt.Errorf("error decoding response: %v", err)
		}

		siteIDMap := make(map[string]string)
		for _, id := range batch {
			siteIDMap[id.Site] = id.ID
		}

		for _, item := range response {
			if item.Code == http.StatusOK {
				results = append(results, ResponseData{
					Site:       item.Body.SiteID,
					ID:         siteIDMap[item.Body.SiteID],
					Price:      item.Body.Price,
					StartTime:  item.Body.StartTime,
					CategoryID: item.Body.CategoryID,
					CurrencyID: item.Body.CurrencyID,
					SellerID:   item.Body.SellerID,
				})
			}
		}

		// log.Printf("Batch processed: %v", batch)
	}

	log.Printf("FetchData completed with results: %v", results)
	return results, nil
}
