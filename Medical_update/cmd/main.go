
package main

import (
    "actualizacion-medica/config"
    "actualizacion-medica/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    r := gin.Default()
    routes.Routes(r)
    r.Run(":4001")
}
