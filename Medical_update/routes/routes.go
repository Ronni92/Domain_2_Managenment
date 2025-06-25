
package routes

import (
    "actualizacion-medica/controllers"
    "github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
    r.POST("/api/actualizaciones", controllers.CrearActualizacion)
}
