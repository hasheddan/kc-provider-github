# Set the shell to bash always
SHELL := /bin/bash

# Options
ORG_NAME=hasheddan
PROVIDER_NAME=kc-provider-github

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/$(PROVIDER_NAME)-controller cmd/provider/main.go

image:
	docker build . -t $(ORG_NAME)/$(PROVIDER_NAME)-controller:v0.0.2 -f cluster/Dockerfile

image-push:
	docker push $(ORG_NAME)/$(PROVIDER_NAME)-controller:v0.0.2

run: generate
	kubectl apply -f package/crds/ -R
	go run cmd/provider/main.go -d

all: image image-push install

generate:
	go generate ./...

lint:
	$(LINT) run

tidy:
	go mod tidy

test:
	go test -v ./...

# Tools

KIND=$(shell which kind)
LINT=$(shell which golangci-lint)

.PHONY: generate tidy lint clean build image all run