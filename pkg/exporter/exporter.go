package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"strings"
	"time"
)

const (
	namespace = "yh"
	platform  = "cmp"
	region    = "beijing"
)

var (
	label = []string{"platform", "region"}

	storageUsage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "storage_usage"),
		"record storage_usage",
		label, nil,
	)
)

type Exporter struct {
	url string
}

// NewExporter returns an initialized Exporter.
func NewExporter(url string) (*Exporter, error) {
	if !strings.Contains(url, "://") {
		url = "http://" + url
	}

	// Init our exporter.
	return &Exporter{
		url: url,
	}, nil
}

// Describe describes all the metrics ever exported by the Consul exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- storageUsage
}

// Collect fetches the stats from configured Consul location and delivers them
// as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	log.Infoln("prometheus pull", time.Now())

	client, err := NewStorage(&Storage{})
	if err != nil {
		log.Errorln(err)
		return
	}

	avg, err := client.GetUsage()
	if err == nil {
		ch <- prometheus.MustNewConstMetric(storageUsage, prometheus.GaugeValue, float64(avg), platform, region)
	} else {
		log.Errorln(err)
	}

	log.Infoln("prometheus pull metrics end ", time.Now())
}
