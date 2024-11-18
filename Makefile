testdir ?= "./test/..."
u_test:
	go test -v $(testdir)

env:
	cp .env.example .env && cp .env.example .env.test

download:
	go mod download

setup:
	env
	download

run:
	go run cmd/main.go

dev:
	./scripts/dev.sh
