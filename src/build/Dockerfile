FROM golang:1.22-bookworm as builder

RUN mkdir "/tmp/build"

WORKDIR /tmp/build

COPY . .

RUN go mod tidy &&  \
    go mod download &&  \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/app/main.go

FROM debian:12-slim

COPY --from=builder /tmp/build/main /usr/local/bin/main

CMD main
