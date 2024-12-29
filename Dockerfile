FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY internal/ ./internal/

COPY . .

RUN ls -R /app

RUN go build -o gopack ./main

