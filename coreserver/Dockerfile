FROM golang:1.8

ENV PROJECT $GOPATH/src/github.com/zyfdegh/fanach/coreserver
WORKDIR $PROJECT

COPY . $PROJECT

RUN	go build -o coreserver && \
	go test $(go list ./... | grep -v /vendor/)

CMD ["./coreserver"]
