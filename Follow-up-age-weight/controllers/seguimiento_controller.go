// controllers/seguimiento_controller.go
package controllers

import (
	"context"
	"net/http"
	"os"
	"seguimiento-peso-edad/config"
	"seguimiento-peso-edad/models"
	"seguimiento-peso-edad/services"
	"time"

	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func RegistrarSeguimiento(c *gin.Context) {
	var s models.Seguimiento
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	estrategia := services.PesoEstrategia{}
	campos := estrategia.Procesar(s)

	estrategia2 := services.EdadEstrategia{}
	for k, v := range estrategia2.Procesar(s) {
		campos[k] = v
	}

	punto := influxdb2.NewPoint(
		"seguimiento",
		map[string]string{"mascota_id": s.MascotaID},
		campos,
		time.Now(),
	)

	writeAPI := config.Cliente.WriteAPIBlocking(os.Getenv("INFLUX_ORG"), os.Getenv("INFLUX_BUCKET"))
	err := writeAPI.WritePoint(context.Background(), punto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Seguimiento registrado"})
}
