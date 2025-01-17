package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"project/api"
	"project/utils"

	"github.com/gin-gonic/gin"
)

var config utils.FileConfig

func init() {
	var err error
	config, err = utils.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	log.Printf("Configuración cargada: %+v", config)
}

func main() {
	// Inicializar servidor Gin
	r := gin.Default()

	// Endpoint para procesar archivos pendientes
	r.POST("/process-files", func(c *gin.Context) {
		log.Println("Solicitud recibida para procesar archivos pendientes")

		// Procesar y enviar archivos
		err := api.ProcessAndSendFiles(config)
		if err != nil {
			log.Printf("Error al procesar archivos: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar archivos"})
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
