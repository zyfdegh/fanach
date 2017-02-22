#!/bin/sh

go get github.com/kataras/iris
go build -o bin/coreserver

if [[ $? -ne 0 ]]; then
	 exit $?
fi

sudo docker build -t zyfdedh/fanach-coreserver:dev .

# docker run --rm -p 18080:8080 zyfdedh/fanach-coreserver:dev
