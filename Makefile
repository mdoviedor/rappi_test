
SHELL := /bin/bash

PACKAGES := $(shell go list ./... | grep -vE "vendor" )

fmt:
	go fmt $(PACKAGES)

test:
	go test -timeout 60s $(PACKAGES)

cover:
	go-acc $(PACKAGES)
	go tool cover -html=coverage.txt -o coverage.html

.PHONY: test  cover fmt
