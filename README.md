```
         ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
Go-prom
```
# About
This is based on the Prometheus docs ([Instrumenting a GO application for Prometheus](https://prometheus.io/docs/guides/go-application/)).

I just take that post as a base project with the idea to create a simple Go application, deploy on a container and then try some Prometheus metrics on a local and kubernetes premises, in the way to make some tests with custom metrics.

This repo will be public, collaborative to anyone who wants to write some code just for fun.

## Getting started


## How to make it run:

### Local Docker standalone

Fork [the Go-prom repo](https://github.com/cosckoya/go-prom) and:

```sh
  $ git clone git@github.com:my-github-user/go-prom.git
  $ cd go-prom
  $ docker build -t my-github-user/go-prom .
  $ docker run -p80:2112 my-github-user/go-prom
```

At this point Prometheus should be serving metrics at **http://localhost/metrics**.

Run on terminal:
```bash
$> curl http://localhost/metrics
```
And you will be able to see some Prometheus metrics


### Local Docker + Prometheus

Now it's time to make thing more complex. Let's try to start a Prometheus+Grafana to see some dashboards with our desired metrics.

First you will need to install [docker-compose](https://github.com/docker/compose), configure it will be fine too :P

Then run the following commands:
```bash
$> docker-compose build
$> docker-compose up
```
At this point you will be able to navitage into [Prometheus](http://localhost:9090) and [Grafana](http://localhost:3000) (grafana login: admin/eatme)

As you could see there is a dashboard with some basic Go metrics!

Ps. To clean this mess you should run the following commands:
```bash
$> docker-compose stop
$> docker-compose rm
$> docker volume rm go-prom_grafana_data go-prom_prometheus_data
```
Enjoy!

### Local Kubernetes (Kind sample)

TBD: This will have some Helm and YAML definitions to deploy a Kubernetes environment (with [Kind](https://kind.sigs.k8s.io))
