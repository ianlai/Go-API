FROM golang:1.16.3-buster

COPY main.go ./

RUN ["go", "run", "main.go"]
