package mediator

import (
	"encoding/json"
	"fmt"
	"log"
	api "project/apis"
	"project/utils"
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
	for _, data := range itemsAPIResponse {
		fmt.Println("Processing item: ", data)
		name, err := m.categoryAPI.FetchData(data.CategoryID)
		if err != nil {
			return "", err
		}
		description, err := m.currencyAPI.FetchData(data.CurrencyID)
		if err != nil {
			return "", err
		}
		nickname, err := m.usersAPI.FetchNumericData(data.SellerID)
		if err != nil {
			return "", err
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
		db_response, err := utils.DoQuery(query, args)
		if err != nil {
			return "", fmt.Errorf("error inserting data into database: %v", err)
		}
		var result []*utils.AffectsRows

		// Deserializar el arreglo de bytes a la estructura AffectsRows
		if err := json.Unmarshal(db_response, &result); err != nil {
			log.Println("Error al deserializar los datos:", err)
			return "", err
		}
	}

	return "", nil
}
