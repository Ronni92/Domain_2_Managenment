// models/seguimiento.go
package models

type Seguimiento struct {
	MascotaID string  `json:"mascota_id"`
	Peso      float64 `json:"peso"`
	Edad      int     `json:"edad"`
}
