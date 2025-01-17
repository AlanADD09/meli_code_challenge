package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	api "project/apis"
	"project/mediator"
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
	r := gin.Default()
	// Endpoint para recibir archivos
	r.POST("/receive-files", func(c *gin.Context) {
		var data []api.SiteID
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		manager := mediator.NewAPIManager(config)
		_, err := manager.Process(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "archivos recibidos"})
	})
	// Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Iniciando servidor en el puerto %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
