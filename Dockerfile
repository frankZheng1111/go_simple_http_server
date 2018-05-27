FROM golang:latest

WORKDIR $GOPATH/src/go_simple_http_server
ADD . $GOPATH/src/go_simple_http_server
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./go_simple_http_server"]

