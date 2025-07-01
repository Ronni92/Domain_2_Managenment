// config/influxdb.go
package config

import (
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var Cliente influxdb2.Client

func ConectarInflux() {
	url := os.Getenv("INFLUX_URL")
	token := os.Getenv("INFLUX_TOKEN")
	Cliente = influxdb2.NewClient(url, token)
	log.Println("✅ Conectado a InfluxDB")
}
