# eso-sops-server
eso-sops-server is a backend for External Secret Operator and expose Mozilla Sops encrpted secrtes

## Getting started

Basic tasks after cloning:

- Build and run locally:

	```sh
	make build
	PORT=8080 ./bin/server
	```

- Run tests:

	```sh
	make test
	```

- Build Docker image:

	```sh
	make docker-build
	```

CI: a GitHub Actions workflow is included at `.github/workflows/ci.yml` that runs tests and builds a Linux binary artifact.

This repository contains a minimal HTTP server under `cmd/server` and a health endpoint at `/healthz`.
