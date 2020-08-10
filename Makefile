APP_NAME := delayed-queue
VERSION ?= $(shell git describe --always)

GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")
PACKAGES ?= $(shell go list ./ ... | grep -v /vendor/)

vet:
	go vet $(PACKAGES)

fmt:
	gofmt -l -s -w $(GOFILES)

deps:
	go mod tidy

build:
	CGO_ENABLED=0 \
	go build \
	-v \
	-o $(APP_NAME)-$(VERSION) .

run:
	go run .

test:
	go test -race -coverprofile coverage.out -v ./...
	go tool cover -html=coverage.out -o coverage_report.html
