all: build

dep:
	go get github.com/kardianos/govendor
	govendor update

build:
	docker build -t zyfdedh/fanach-coreserver .

run:
	docker run --rm -p 8080:8080 zyfdedh/fanach-coreserver

push: build
	docker push zyfdedh/fanach-coreserver .
