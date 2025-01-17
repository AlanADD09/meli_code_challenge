package api

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	// "project/mediator"
	"strconv"
)

type UserAPI struct {
	// mediator     mediator.Mediator
	BaseURL     string
	BearerToken string
}

func (api *UserAPI) FetchMultiData(input []SiteID) ([]ResponseData, error) {
	return nil, nil
}

func (api *UserAPI) FetchData(input string) (string, error) {
	return "", nil
}

func (api *UserAPI) FetchNumericData(input int64) (string, error) {
	fmt.Println("Fetching user data")
	var result string
	// Construir la URL para el batch
	url := fmt.Sprintf("%s/users/%s", api.BaseURL, strconv.FormatInt(input, 10))

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
		Nickname string `json:"nickname"`
	}

	// fmt.Println("Response: ", resp)

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error al decodificar respuesta: %v", err)
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	result = response.Nickname

	fmt.Println("Result: ", result)

	// log.Printf("User processed for ID: %s", data.UserID)

	// log.Printf("FetchData completed with results for user")
	return result, nil
}
