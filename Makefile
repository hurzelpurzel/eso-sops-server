BINARY=bin/server

.PHONY: all build run test docker-build clean

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY) ./cmd/server

run:
	# Run locally (port can be set with PORT env)
	PORT=8080 $(BINARY)

test:
	go test ./... -v

docker-build:
	docker build -t hurzelpurzel/eso-sops-server:latest .

clean:
	rm -f $(BINARY)
