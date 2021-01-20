package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
)

//http://vearne.cc/archives/39380

func init() {
	prometheus.MustRegister(cpuTemp)
}

func Run() {

	go func() {
		for {
			rand.NewSource(time.Now().Unix())
			cpuTemp.Set(rand.Float64() * 100)
			time.Sleep(time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
