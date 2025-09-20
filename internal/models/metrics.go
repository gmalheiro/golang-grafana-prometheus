package models

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	devices prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Name:      "myApp",
			Namespace: "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
	}

	reg.MustRegister(m.devices)

	return m
}

func (m *metrics) Devices() prometheus.Gauge {
	return m.devices
}
