FROM golang AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY main.go .
RUN go build -o main .

# ---

FROM golang:alpine AS app

WORKDIR /app

COPY --from=build /build/main .
COPY template.html .
COPY img/ /app/img

EXPOSE 8080

CMD ["/app/main"]
