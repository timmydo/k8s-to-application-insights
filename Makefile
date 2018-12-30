HASH := $(shell git rev-parse --short HEAD)

.PHONY: all
all: build

.PHONY: build
build:
	go build -o k8s-to-ai

.PHONY: version
version:
	echo $(HASH)

.PHONY: docker
docker:
	docker build -t timmydo/k8s-to-application-insights:git-$(HASH) .

.PHONY: dev
dev:
	docker run --rm -it -e GO111MODULE=on --workdir /go/src/github.com/timmydo/k8s-to-application-insights --volume $(CURDIR):/go/src/github.com/timmydo/k8s-to-application-insights quay.io/deis/go-dev:latest

.PHONY: run
run: build
	./k8s-to-ai -port 8080