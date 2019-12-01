FROM golang:latest

ADD ../../ ./

RUN apt-get update
RUN apt-get install -y mysql-client
RUN apt-get install -y vim
RUN apt-get install -y git
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose
RUN git clone https://github.com/k-ueki/Agpluss.git

WORKDIR /go/src/agpluss

RUN go build
