FROM golang:1.22.4-alpine3.20

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...
