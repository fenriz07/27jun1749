#BaseImage arg
FROM golang:1.15-alpine as builder

COPY . /app
WORKDIR /app
RUN go mod vendor

RUN go build -v -o server

EXPOSE 8080

CMD  /app/server

