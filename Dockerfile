FROM golang:1.13

WORKDIR /app

COPY * ./

RUN go mod download