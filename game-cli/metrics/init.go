package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	delta struct {
		delta    float64
		lastTick time.Time
		hist     prometheus.Histogram
	}

	Delta interface {
		Tick() float64
	}
)

var (
	globalMetrics struct {
		delta *delta
	}
)

func RunMetricsHTTP() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("localhost:2112", nil)
}

func GetDelta() Delta {
	return globalMetrics.delta
}

func (d *delta) Tick() float64 {
	now := time.Now()
	d.delta = now.Sub(d.lastTick).Seconds()
	d.lastTick = now
	d.hist.Observe(d.delta)
	return d.delta
}

func init() {
	globalMetrics.delta = &delta{
		delta:    0,
		lastTick: time.Now(),
		hist: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:      "fps",
			Subsystem: "gamecli",
			Buckets: []float64{
				0.005, 0.01, 0.1, 0.2, 0.5, 0.7, 1,
			},
		}),
	}

}
