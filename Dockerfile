FROM golang:1.26-alpine

RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED=1

RUN go install github.com/air-verse/air@latest

WORKDIR /var/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 3000

ENTRYPOINT [ "air" ]