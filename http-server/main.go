package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "github.com/urfave/negroni"
)

type KittensPageData struct {
    KittensPageTitle string
}

var (
    httpServerSecondsUp = promauto.NewCounter(prometheus.CounterOpts{
        Namespace: "http",
        Name: "kittens_server_up",
        Help: "Number of seconds this HTTP server has been serving kittens' webpage",
    })

    requestDurations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
        Namespace: "http",
        Name: "kittens_request_seconds",
        Help: "The latency of the HTTP requests.",
        Buckets: prometheus.DefBuckets,
    }, []string{"name", "path"})
)

func recordMetrics() {
    go func() {
        for {
            httpServerSecondsUp.Inc()
            time.Sleep(1 * time.Second)
        }
    }()
}

func requestDuration(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    var startedAt = time.Now()
    next(rw, r)
    requestDurations.WithLabelValues("kittens_request_seconds", r.URL.Path).Observe(time.Since(startedAt).Seconds())
}


func main() {
    mux := http.NewServeMux()
    tmpl := template.Must(template.ParseFiles("template.html"))

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := KittensPageData{
            KittensPageTitle: "My kittens webpage!!",
        }
        tmpl.Execute(w, data)
    })

    fs := http.FileServer(http.Dir("img/"))
    mux.Handle("/img/", http.StripPrefix("/img/", fs))

    mux.Handle("/metrics", promhttp.Handler())

    n := negroni.Classic() // Includes some default middlewares
    n.UseHandler(mux)

    recordMetrics()

    prometheus.Register(requestDurations)
    n.Use(negroni.HandlerFunc(requestDuration))

    http.ListenAndServe(":8080", n)
}
