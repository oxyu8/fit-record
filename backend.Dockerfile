FROM golang:latest

RUN apt-get update && apt-get install -y git
WORKDIR /go/src/fit_record_server