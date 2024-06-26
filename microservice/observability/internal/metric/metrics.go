package metric

import "github.com/prometheus/client_golang/prometheus"

var TestCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace:   "myapp",
	Subsystem:   "api",
	Name:        "request_counter",
	Help:        "Counter for requests",
	ConstLabels: nil,
})
