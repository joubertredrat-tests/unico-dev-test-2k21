FROM golang:1.16-alpine

RUN apk add gcc g++ make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN touch .env

ENTRYPOINT [ "/bin/sh", "/app/entrypoint.sh" ]