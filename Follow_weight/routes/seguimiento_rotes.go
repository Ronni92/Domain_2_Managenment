// routes/seguimiento_routes.go
package routes

import (
	"seguimiento-peso-edad/controllers"

	"github.com/gin-gonic/gin"
)

func RegistrarRutas(r *gin.Engine) {
	r.POST("/api/seguimiento", controllers.RegistrarSeguimiento)
}
