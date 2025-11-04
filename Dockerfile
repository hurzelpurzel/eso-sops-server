# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /workspace

# Install git to fetch modules if needed
RUN apk add --no-cache git

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the Go source (relies on .dockerignore to filter)
COPY . .



RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

###################### Final stage ###########################
FROM alpine:3.19
RUN apk add --no-cache ca-certificates curl gnupg
# Install age (used by SOPS)
RUN curl -L https://github.com/FiloSottile/age/releases/download/v1.1.1/age-v1.1.1-linux-amd64.tar.gz \
    | tar -xz && \
    mv age/age age/age-keygen /usr/local/bin/ && \
    rm -rf age

# Install SOPS
RUN curl -L https://github.com/getsops/sops/releases/download/v3.8.1/sops-v3.8.1.linux \
    -o /usr/local/bin/sops && \
    chmod +x /usr/local/bin/sops

# Optional: verify installation
#RUN sops --version && age --version

VOLUME /config

COPY --from=builder /out/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
