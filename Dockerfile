FROM golang:1.17

ADD . /app

WORKDIR /app

RUN go mod download
RUN go build -o main .

CMD ["/app/main"]
