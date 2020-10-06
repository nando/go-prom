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
