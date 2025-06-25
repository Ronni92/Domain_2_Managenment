
package models

type ActualizacionMedica struct {
    ID           string `json:"id" bson:"_id,omitempty"`
    MascotaID    string `json:"mascota_id" bson:"mascota_id"`
    Diagnostico  string `json:"diagnostico" bson:"diagnostico"`
    Tratamiento  string `json:"tratamiento" bson:"tratamiento"`
    Fecha        string `json:"fecha" bson:"fecha"`
}
