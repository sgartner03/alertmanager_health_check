package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)


type Metrics struct {
	CounterVec *prometheus.CounterVec
}

func NewMetrics(name string, help string, label string) Metrics {
	var metrics Metrics
	metrics.SetCounterVec(name, help, label)
	return metrics
}

func (metrics *Metrics) SetCounterVec(name string, help string, label string)  {
	metrics.CounterVec = promauto.NewCounterVec(prometheus.CounterOpts{
                Name: name, // "alertmanager_status",
                Help: help,}, //"The status of the alertmanager",},
                []string{label}) //"gepardec_cluster"})
}

func (metrics *Metrics) Increment(label string) {
	metrics.CounterVec.WithLabelValues(label).Inc()
}

func (metrics *Metrics) IncrementSequence(labels []string) {
	for _, label := range labels {
		metrics.Increment(label)
	}
}
