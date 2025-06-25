// main.go
package main

import (
	"log"
	"seguimiento-peso-edad/config"
	"seguimiento-peso-edad/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	config.ConectarInflux()

	r := gin.Default()
	routes.RegistrarRutas(r)
	r.Run(":3004")
}
