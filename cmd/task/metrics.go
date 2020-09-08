package main

import (
	"runtime"

	"github.com/Djarvur/allcups-itrally-2020-task/pkg/def"
	"github.com/prometheus/client_golang/prometheus"
)

func initMetrics(reg *prometheus.Registry, namespace string) {
	reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	reg.MustRegister(prometheus.NewGoCollector())

	version := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "build_info",
			Help:      "A metric with a constant '1' value labeled by build-time details.",
		},
		[]string{"version", "goversion"},
	)
	reg.MustRegister(version)

	version.With(prometheus.Labels{
		"version":   def.Version(),
		"goversion": runtime.Version(),
	}).Set(1)
}
