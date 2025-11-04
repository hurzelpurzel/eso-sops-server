# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /src

# Install git to fetch modules if needed
RUN apk add --no-cache git

COPY go.mod go.sum* ./
RUN go env -w GOPROXY=https://proxy.golang.org || true
RUN go mod download || true

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

# Final stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
