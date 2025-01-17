package api_factory

import (
	"database/sql"
	"fmt"
)

// Repository define la interfaz para interactuar con la base de datos
type Repository interface {
	SaveData(data map[string]interface{}) error
}

// PostgreSQLRepository implementa Repository para PostgreSQL
type PostgreSQLRepository struct {
	DB *sql.DB
}

func (repo PostgreSQLRepository) SaveData(data map[string]interface{}) error {
	query := "INSERT INTO items (id, name, price, category) VALUES ($1, $2, $3, $4)"
	_, err := repo.DB.Exec(query, data["id"], data["name"], data["price"], data["category"])
	if err != nil {
		return fmt.Errorf("error al guardar datos en la base de datos: %w", err)
	}
	return nil
}
