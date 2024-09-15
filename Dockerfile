FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make tidy
RUN make build

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/bin/app /app/

CMD ["/app/app"]