package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project/api_factory"
	"project/file_processor"
	"project/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	config, err := utils.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	log.Printf("Configuración cargada: %+v", config)

	// Inicializar servidor Gin
	r := gin.Default()

	// Endpoint para procesar archivos pendientes
	r.POST("/process-files", func(c *gin.Context) {
		log.Println("Solicitud recibida para procesar archivos pendientes")

		// Inicializar cliente de la API
		apiClient, err := api_factory.GetAPIClient("mercadolibre", "https://api.mercadolibre.com", config.BearerToken)
		if err != nil {
			log.Fatalf("Error al inicializar cliente API: %v", err)
		}

		if err := file_processor.ProcessPendingFiles(config, apiClient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Archivos procesados con éxito"})
	})

	// Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Iniciando servidor en el puerto %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
