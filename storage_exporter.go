package main

import (
	"flag"
	"git.dev.yonghui.cn/mqiqe/storage_exporter/pkg/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"net/http"
)

var (
	listenAddress string
	metricsPath   string
	url           string
)

func init() {
	flag.StringVar(&listenAddress, "listen-address", ":8096", "The address to listen on for HTTP requests.")
	flag.StringVar(&metricsPath, "path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&url, "url", "localhost:8096", "exporter url")
	flag.Parse()
}
func main() {
	log.Infoln("Starting storage_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	exporter, err := exporter.NewExporter(url)
	if err != nil {
		log.Fatalln(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Storage Exporter</title></head>
			<body>
			<h1>Node Exporter</h1>
			<p><a href="` + metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
