From golang:latest

ADD . /go/src/github.com/eoengineer/test-api

RUN go install github.com/eoengineer/test-api

ENTRYPOINT /go/bin/test-api

EXPOSE 8080
