// services/edad_strategy.go
package services

import "seguimiento-peso-edad/models"

type EdadEstrategia struct{}

func (e EdadEstrategia) Procesar(s models.Seguimiento) map[string]interface{} {
	return map[string]interface{}{
		"edad": s.Edad,
	}
}
