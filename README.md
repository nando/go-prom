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
=======
# Go-prom
## Getting started

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

TBD: This will have some nice docker-compose file with a local container image and a Prometheus server, maybe a Grafana too :D

### Local Kubernetes (Kind sample)

TBD: This will have some Helm and YAML definitions to deploy a Kubernetes environment (with [Kind](https://kind.sigs.k8s.io))