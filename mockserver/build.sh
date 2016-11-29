#!/bin/sh

go build -o bin/mockserver

sudo docker build -t zyfdedh/fanach-mockserver .

# docker run --rm -p 18080:8080 zyfdedh/fanach-mockserver
