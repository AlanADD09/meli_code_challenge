package api

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	// "project/mediator"
)

type CurrencyAPI struct {
	// mediator     mediator.Mediator
	BaseURL     string
	BearerToken string
}

func (api *CurrencyAPI) FetchMultiData(input []SiteID) ([]ResponseData, error) {
	return nil, nil
}

func (api *CurrencyAPI) FetchNumericData(input int64) (string, error) {
	return "", nil
}

func (api *CurrencyAPI) FetchData(input string) (string, error) {
	fmt.Println("Fetching currency data")
	var result string
	// Construir la URL para el batch
	url := fmt.Sprintf("%s/currencies/%s", api.BaseURL, input)

	fmt.Println("URL: ", url)

	// Crear la solicitud GET con el Bearer Token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error al crear solicitud HTTP: %v", err)
		return "", fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.BearerToken))

	// Realizar la solicitud
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error al realizar solicitud a la API: %v", err)
		return "", fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error en respuesta de la API: %d", resp.StatusCode)
		return "", fmt.Errorf("API response error: %d", resp.StatusCode)
	}

	// Decodificar la respuesta
	var response struct {
		Description string `json:"description"`
	}

	// fmt.Println("Response: ", response)

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error al decodificar respuesta: %v", err)
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	result = response.Description

	fmt.Println("Result: ", result)

	// log.Printf("Currency processed for ID: %s", data.CurrencyID)

	// log.Printf("FetchData completed with results for currency")
	return result, nil
}
