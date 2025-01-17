package mediator

import (
	"encoding/json"
	"fmt"
	"log"
	api "project/apis"
	"project/utils"
	"sync"
)

type Mediator interface {
	process(inputs []string) (string, error)
}

type APIManager struct {
	itemsAPI    api.API
	categoryAPI api.API
	currencyAPI api.API
	usersAPI    api.API
}

func NewAPIManager(config utils.FileConfig) *APIManager {
	return &APIManager{
		itemsAPI: &api.ItemsAPI{
			BaseURL:     config.MeliURL,
			BearerToken: config.BearerToken,
		},
		categoryAPI: &api.CategoryAPI{
			BaseURL:     config.MeliURL,
			BearerToken: config.BearerToken,
		},
		currencyAPI: &api.CurrencyAPI{
			BaseURL:     config.MeliURL,
			BearerToken: config.BearerToken,
		},
		usersAPI: &api.UserAPI{
			BaseURL:     config.MeliURL,
			BearerToken: config.BearerToken,
		},
	}
}

func (m *APIManager) Process(inputs []api.SiteID) (string, error) {
	query := "INSERT INTO tbl_data (site_id, id, price, date_created, category_id, currency_id, seller_id, Name, Description, Nickname, error) VALUES (@site_id, @id, @price, @date_created, @category_id, @currency_id, @seller_id, @Name, @Description, @Nickname, @error)"

	itemsAPIResponse, err := m.itemsAPI.FetchMultiData(inputs)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	errChan := make(chan error) // Canal sin buffer para recibir errores de goroutines
	done := make(chan struct{}) // Canal para saber cuándo cerrar `errChan`

	// Goroutine separada para cerrar errChan cuando todas las goroutines terminen
	go func() {
		wg.Wait()
		close(errChan) // Se cierra SOLO cuando todas las goroutines terminan
		close(done)    // Notificamos que el proceso finalizó
	}()

	for _, d := range itemsAPIResponse {
		wg.Add(1)
		data := d // Capturar la variable dentro del bucle

		go func(data api.ResponseData) {
			defer wg.Done()

			fmt.Println("Processing item:", data)
			var name, description, nickname string

			if data.CategoryID != "" {
				name, err = m.categoryAPI.FetchData(data.CategoryID)
				if err != nil {
					select {
					case errChan <- fmt.Errorf("error in category API: %v", err):
					default: // Evita enviar al canal si ya se cerró
					}
					return
				}
			}

			if data.CurrencyID != "" {
				description, err = m.currencyAPI.FetchData(data.CurrencyID)
				if err != nil {
					select {
					case errChan <- fmt.Errorf("error in currency API: %v", err):
					default:
					}
					return
				}
			}

			if data.SellerID != 0 {
				nickname, err = m.usersAPI.FetchNumericData(data.SellerID)
				if err != nil {
					select {
					case errChan <- fmt.Errorf("error in users API: %v", err):
					default:
					}
					return
				}
			}

			data.Name = name
			data.Description = description
			data.Nickname = nickname

			args := []utils.SqlArgs{
				{Name: "site_id", Value: data.Site},
				{Name: "id", Value: data.ID},
				{Name: "price", Value: data.Price},
				{Name: "date_created", Value: data.StartTime},
				{Name: "category_id", Value: data.CategoryID},
				{Name: "currency_id", Value: data.CurrencyID},
				{Name: "seller_id", Value: data.SellerID},
				{Name: "Name", Value: data.Name},
				{Name: "Description", Value: data.Description},
				{Name: "Nickname", Value: data.Nickname},
				{Name: "error", Value: data.Error},
			}

			dbResponse, err := utils.DoQuery(query, args)
			if err != nil {
				select {
				case errChan <- fmt.Errorf("error inserting data into database: %v", err):
				default:
				}
				return
			}

			var result []*utils.AffectsRows
			if err := json.Unmarshal(dbResponse, &result); err != nil {
				log.Println("Error al deserializar los datos:", err)
				select {
				case errChan <- err:
				default:
				}
				return
			}
		}(data)
	}

	// Verificamos si hay errores
	select {
	case err := <-errChan:
		return "", err // Retorna el primer error que aparezca
	case <-done:
		// No hay errores, retornamos éxito
		return "", nil
	}
}
