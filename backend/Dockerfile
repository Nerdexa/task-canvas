FROM golang:1.22.6-alpine3.19 as builder

WORKDIR /backend

RUN apk update && apk add git

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]