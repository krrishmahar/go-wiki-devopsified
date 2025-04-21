# Stage 1: Build the Go binary
FROM golang:1.24.2 AS base

WORKDIR /app

# Copy dependencies first for layer caching
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary for Linux
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Use minimal distroless image
FROM gcr.io/distroless/base

WORKDIR /

# Copy binary and required folders from base
COPY --from=base /app/main /main
COPY --from=base /app/tmpl /tmpl
COPY --from=base /app/data /data

EXPOSE 8080
USER nonroot

CMD ["./main"]
