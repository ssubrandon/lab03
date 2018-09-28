FROM golang:1.8-alpine

RUN mkdir ~/go

ENV GOPATH /go

WORKDIR /go

COPY * /go/src/

RUN go get -t app
RUN go install app

CMD ["./bin/app"]
