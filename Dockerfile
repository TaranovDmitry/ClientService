FROM golang:alpine

WORKDIR /clientService

COPY communications ./
COPY config ./
COPY handlers ./
COPY services ./
COPY go.sum ./
COPY go.mod ./
COPY main.go ./

RUN go build -o clientService .

CMD ["/clientService/clientService"]
