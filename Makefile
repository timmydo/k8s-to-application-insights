

.PHONY: all
all: build

.PHONY: build
build:
	go build -o k8s-to-ai.exe

.PHONY: run
run: build
	./k8s-to-ai.exe -port 8080