package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// 		Example
	//stateRandomCounter = prometheus.NewCounter(prometheus.CounterOpts{
	//	Name: "factomd_state_randomcounter_total",
	//	Help: "Just a basic counter that can only go up",
	//})

	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "counter_metric",
		Help: "Just a basic counter that can only go up",
	})

	gauge = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gauge_metric",
		Help: "Just a basic guage that can up or down",
	},
		[]string{"label1", "label2"})

	histo = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "histogram_metric",
		Help:    "A histogram",
		Buckets: prometheus.ExponentialBuckets(1, 2, 10),
	},
		[]string{"action", "address_family", "error"},
	)
)

var registered = false

// RegisterPrometheus registers the variables to be exposed. This can only be run once, hence the
// boolean flag to prevent panics if launched more than once. This is called in NetStart
func RegisterPrometheus() {
	if registered {
		return
	}
	registered = true
	// 		Exmaple Cont.
	// prometheus.MustRegister(stateRandomCounter)

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
}
