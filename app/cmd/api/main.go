package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gmalheiro/golang-grafana-prometheus/internal/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var dvs []models.Device
var prometheusPort string = ":8081"
var devicesPort string = ":8080"
var version string

func init() {
	version = "0.0.1"
	dvs = []models.Device{
		{ID: 1, Mac: "a1b2", Firmware: "2.1.6"},
		{ID: 2, Mac: "a2b3", Firmware: "2.1.6"},
		{ID: 3, Mac: "a3b4", Firmware: "2.1.6"},
	}
}

func main() {
	reg := prometheus.NewRegistry()
	m := models.NewMetrics(reg)

	m.Devices().Set(float64(len(dvs)))
	m.Info().With(prometheus.Labels{"version": version})

	dmux := http.NewServeMux()
	dmux.HandleFunc("/devices", getDevices)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pmux := http.NewServeMux()
	pmux.Handle("/metrics", promHandler)

	go func() {
		fmt.Printf("\nListening devices on port %s", devicesPort)
		if err := http.ListenAndServe(devicesPort, dmux); err != nil {
			log.Fatal()
		}

	}()

	go func() {
		fmt.Printf("Listening prometheus on port %s", prometheusPort)
		if err := http.ListenAndServe(prometheusPort, pmux); err != nil {
			log.Fatal()
		}
	}()

	select {}
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
