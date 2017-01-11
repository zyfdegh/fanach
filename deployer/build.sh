#!/bin/sh

go build -o bin/deployer
cp -r static bin/

if [[ $? -ne 0 ]]; then
	 exit $?
fi

sudo docker build -t zyfdedh/fanach-deployer:dev .

# docker run --rm -p 18080:8080 zyfdedh/fanach-deployer:dev
