// process_files.go
package file_processor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"project/api_factory"
	"project/utils"
)

// // ProcessPendingFiles procesa los archivos pendientes en el directorio configurado
// func ProcessPendingFiles(config utils.FileConfig) error {
// 	log.Printf("Procesando archivos en el directorio: %s", config.Directory)

// 	// Listar archivos en el directorio
// 	files, err := os.ReadDir(config.Directory)
// 	if err != nil {
// 		return fmt.Errorf("error al leer el directorio: %w", err)
// 	}

// 	// Mapear formatos a extensiones válidas
// 	validExtensions := map[string]string{
// 		"csv":       ".csv",
// 		"jsonlines": ".jsonl",
// 	}

// 	expectedExtension, ok := validExtensions[config.Format]
// 	if !ok {
// 		return fmt.Errorf("formato configurado no soportado: %s", config.Format)
// 	}

// 	for _, file := range files {
// 		if file.IsDir() {
// 			continue
// 		}

// 		// Verificar la extensión del archivo
// 		filePath := fmt.Sprintf("%s/%s", config.Directory, file.Name())
// 		if filepath.Ext(file.Name()) != expectedExtension {
// 			log.Printf("Saltando archivo %s: no coincide con el formato configurado (%s)", file.Name(), config.Format)
// 			continue
// 		}

// 		log.Printf("Procesando archivo: %s", filePath)

// 		// Obtener el lector adecuado basado en la configuración
// 		reader, err := GetReader(filePath, config)
// 		if err != nil {
// 			log.Printf("Error al obtener el lector para el archivo %s: %v", filePath, err)
// 			continue
// 		}

// 		// Abrir el archivo
// 		f, err := os.Open(filePath)
// 		if err != nil {
// 			log.Printf("Error al abrir el archivo %s: %v", filePath, err)
// 			continue
// 		}

// 		// Leer los datos del archivo
// 		records, err := reader.Read(f)
// 		f.Close()
// 		if err != nil {
// 			log.Printf("Error al procesar el archivo %s: %v", filePath, err)
// 			continue
// 		}

// 		log.Printf("Archivo procesado con éxito: %s, Registros: %d", filePath, len(records))
// 	}

// 	log.Println("Procesamiento de archivos completado")
// 	return nil
// }

// ProcessPendingFiles procesa los archivos pendientes en el directorio configurado
func ProcessPendingFiles(config utils.FileConfig, apiClient api_factory.APIClient) error {
	log.Printf("Procesando archivos en el directorio: %s", config.Directory)

	// Listar archivos en el directorio
	files, err := os.ReadDir(config.Directory)
	if err != nil {
		return fmt.Errorf("error al leer el directorio: %w", err)
	}

	// Mapear formatos a extensiones válidas
	validExtensions := map[string]string{
		"csv":       ".csv",
		"jsonlines": ".jsonl",
	}

	expectedExtension, ok := validExtensions[config.Format]
	if !ok {
		return fmt.Errorf("formato configurado no soportado: %s", config.Format)
	}

	var allItemIDs []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Verificar la extensión del archivo
		filePath := fmt.Sprintf("%s/%s", config.Directory, file.Name())
		if filepath.Ext(file.Name()) != expectedExtension {
			log.Printf("Saltando archivo %s: no coincide con el formato configurado (%s)", file.Name(), config.Format)
			continue
		}

		log.Printf("Procesando archivo: %s", filePath)

		// Obtener el lector adecuado basado en la configuración
		reader, err := GetReader(filePath, config)
		if err != nil {
			log.Printf("Error al obtener el lector para el archivo %s: %v", filePath, err)
			continue
		}

		// Abrir el archivo
		f, err := os.Open(filePath)
		if err != nil {
			log.Printf("Error al abrir el archivo %s: %v", filePath, err)
			continue
		}

		// Leer los datos del archivo
		records, err := reader.Read(f)
		f.Close()
		if err != nil {
			log.Printf("Error al procesar el archivo %s: %v", filePath, err)
			continue
		}

		log.Printf("Archivo procesado con éxito: %s, Registros: %d", filePath, len(records))

		for _, record := range records {
			itemID := record["id"]
			allItemIDs = append(allItemIDs, itemID)
		}
	}

	// Enviar todos los itemIDs a FetchData en lotes
	apiData, err := apiClient.FetchData(allItemIDs)
	if err != nil {
		log.Printf("Error al consultar la API: %v", err)
		return err
	}

	// Procesar y guardar apiData
	for _, data := range apiData["data"].([]api_factory.ProductData) {
		mergedData := map[string]interface{}{
			// "id":       data.ID,
			// "name":     data.Title,
			"price":    data.Price,
			"category": data.CategoryID,
		}

		// if repo == nil {
		// 	log.Printf("Repository is not initialized")
		// 	return fmt.Errorf("repository is not initialized")
		// }

		// if err := repo.SaveData(mergedData); err != nil {
		// 	log.Printf("Error al guardar los datos en la base de datos: %v", err)
		// 	continue
		// }

		log.Printf("Registro procesado y almacenado: %v", mergedData)
	}

	log.Println("Procesamiento de archivos completado")
	return nil
}
