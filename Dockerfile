FROM golang:1.25-alpine AS builder

WORKDIR /var/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /var/app/bin/api ./cmd/api
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /var/app/bin/migrate ./cmd/migrate

FROM alpine:3.22

RUN apk add --no-cache ca-certificates && \
    adduser -D -H -u 10001 appuser

WORKDIR /var/app

COPY --from=builder /var/app/bin/api ./api
COPY --from=builder /var/app/bin/migrate ./migrate
COPY --from=builder /var/app/migrations ./migrations

USER appuser

EXPOSE 3000

CMD [ "./api" ]