package utils

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	uptime = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "node",
			Name:      "uptime",
			Help:      "uptime of the system in seconds",
		},
	)
	load = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "node",
			Name:      "load",
			Help:      "load average, with a label for 1m, 5m & 15m",
		},
		[]string{
			"duration",
		})
)

func SetUptime() {
	go func() {
		for {
			uptimeStr, _ := GetUptime()
			uptimefloat, err := strconv.ParseFloat(uptimeStr, 2)
			if err != nil {
				return
			}
			uptime.Set(uptimefloat)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func SetLoad() {
	go func() {
		for {
			loadStr, _ := GetLoad()
			loadMap := make(map[string]float64)
			err := json.Unmarshal([]byte(loadStr), &loadMap)
			if err != nil {
				return
			}

			load.With(prometheus.Labels{"duration": "1m"}).Set(loadMap["1m"])
			load.With(prometheus.Labels{"duration": "5m"}).Set(loadMap["5m"])
			load.With(prometheus.Labels{"duration": "15m"}).Set(loadMap["15m"])
			time.Sleep(1 * time.Second)
		}
	}()
}
