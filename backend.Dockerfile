FROM golang:1.24

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/cassist-service ./cmd/main.go

# Final image for production
FROM debian:bullseye-slim

WORKDIR /cassist

COPY --from=builder /usr/local/bin/cassist-service .

CMD ["./cassist-service"]
