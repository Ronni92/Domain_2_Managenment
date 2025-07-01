// services/peso_strategy.go
package services

import "seguimiento-peso-edad/models"

type PesoEstrategia struct{}

func (p PesoEstrategia) Procesar(s models.Seguimiento) map[string]interface{} {
	return map[string]interface{}{
		"peso": s.Peso,
	}
}
