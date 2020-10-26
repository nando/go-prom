package main

import (
    "net/http"
    "html/template"
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
)

func main() {
    tmpl := template.Must(template.ParseFiles("template.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := KittensPageData{
            KittensPageTitle: "My kittens webpage!!",
        }
        tmpl.Execute(w, data)
    })

    fs := http.FileServer(http.Dir("img/"))
    http.Handle("/img/", http.StripPrefix("/img/", fs))

    recordMetrics()
    http.Handle("/metrics", promhttp.Handler())

    http.ListenAndServe(":8080", nil)
}
