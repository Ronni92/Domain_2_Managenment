// services/strategy.go
package services

import "seguimiento-peso-edad/models"

type Estrategia interface {
	Procesar(s models.Seguimiento) map[string]interface{}
}
