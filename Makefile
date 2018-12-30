

.PHONY: all
all: build

.PHONY: build
build:
	go build -o k8s-to-ai

.PHONY: docker
docker:
	docker build -t timmydo/k8s-to-application-insights:latest .

.PHONY: dev
dev:
	docker run --rm -it -e GO111MODULE=on --workdir /go/src/github.com/timmydo/k8s-to-application-insights --volume $(CURDIR):/go/src/github.com/timmydo/k8s-to-application-insights quay.io/deis/go-dev:latest

.PHONY: run
run: build
	./k8s-to-ai -port 8080