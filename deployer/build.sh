#!/bin/sh

go get golang.org/x/crypto/ssh
go build -o bin/deployer
cp -r static bin/

sudo docker build -t zyfdedh/fanach-deployer .

# docker run --rm -p 18080:8080 zyfdedh/fanach-deployer
