FROM golang:1.12.0-alpine3.9

RUN mkdir /microfrontends

ADD . /microfrontends

WORKDIR /microfrontends

RUN go build -o .bin/main

CMD ["/microfrontends/main"]