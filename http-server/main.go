package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type KittensPageData struct {
	KittensPageTitle string
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kittens_http_server_seconds_up",
		Help: "Number of seconds the HTTP server has been serving kittens' webpage",
	})
	requestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "kittens_http_request_size",
			Help: "Size of requests made to the http server",
		},
		[]string{"code", "method"},
	)
)

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))

	http.HandleFunc("/", promhttp.InstrumentHandlerRequestSize(requestSize, mainPageHandler(tmpl)))

	fs := http.FileServer(http.Dir("img/"))
	http.Handle("/img/", promhttp.InstrumentHandlerRequestSize(requestSize, http.StripPrefix("/img/", fs)))

	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}

func mainPageHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := KittensPageData{
			KittensPageTitle: "My kittens webpage!!",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Fatalf("Failed to execute template: %s", err)
		}
	}
}
