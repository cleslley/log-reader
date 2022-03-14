FROM golang:1.17.2

WORKDIR /usr/local/go/src/log-reader/src

COPY . /usr/local/go/src/log-reader

RUN cd /usr/local/go/src/log-reader/src

RUN go mod tidy

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /logreader

CMD ["/logreader"]
