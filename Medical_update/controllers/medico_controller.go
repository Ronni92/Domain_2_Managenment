
package controllers

import (
    "actualizacion-medica/config"
    "actualizacion-medica/models"
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func CrearActualizacion(c *gin.Context) {
    var nueva models.ActualizacionMedica
    if err := c.ShouldBindJSON(&nueva); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := config.DB.Collection("actualizaciones")
    res, err := collection.InsertOne(ctx, nueva)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"mensaje": "Actualización médica registrada", "id": res.InsertedID})
}
