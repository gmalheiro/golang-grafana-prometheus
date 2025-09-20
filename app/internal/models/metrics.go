package models

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	devices prometheus.Gauge
	info    *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "connected_devices",
			Name:      "myApp",
			Help:      "Number of currently connected devices.",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myApp",
			Name:      "info",
			Help:      "Information about the My App environment",
		}, []string{"version"}),
	}

	reg.MustRegister(m.devices, m.info)

	return m
}

func (m *metrics) Devices() prometheus.Gauge {
	return m.devices
}
func (m *metrics) Info() *prometheus.GaugeVec {
	return m.info
}
