# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build static binary, disables cgo so no glibc dependency
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy static binary
COPY --from=builder /app/server .

# Copy conf folder (with deploy.yml) so app can find config at runtime
COPY --from=builder /app/conf ./conf

EXPOSE 8080

CMD ["./server"]
