all: build

# Build, test, run in local

dep:
	go get github.com/kardianos/govendor
	govendor update

build-local:
	go build -o bin/coreserver

run-local:
	./bin/coreserver

test-local:
	go test ./...

# Build, test, run in Docker container

build:
	docker build -t zyfdedh/fanach-coreserver .

run:
	docker run --rm -p 8080:8080 zyfdedh/fanach-coreserver

push: build
	docker push zyfdedh/fanach-coreserver .
