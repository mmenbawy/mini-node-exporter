package utils

import (
	"encoding/json"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	uptime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "node",
			Name:      "uptime",
			Help:      "uptime of the system in seconds",
		},
	)
	load = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "node",
			Name:      "load",
			Help:      "load average, with a label for 1m, 5m & 15m",
		},
		[]string{
			"duration",
		})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(uptime)
	prometheus.MustRegister(load)
}

func setUptime() {
	uptimeStr, _ := getUptime()
	uptimefloat, err := strconv.ParseFloat(uptimeStr, 2)
	if err != nil {
		return
	}
	uptime.Set(uptimefloat)
}

func setLoad() {
	loadStr, _ := getLoad()
	loadMap := make(map[string]float64)
	err := json.Unmarshal([]byte(loadStr), &loadMap)
	if err != nil {
		return
	}

	load.With(prometheus.Labels{"duration": "1m"}).Set(loadMap["1m"])
	load.With(prometheus.Labels{"duration": "5m"}).Set(loadMap["5m"])
	load.With(prometheus.Labels{"duration": "15m"}).Set(loadMap["15m"])
}
