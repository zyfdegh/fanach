#!/bin/sh

go build -o bin/dctl

if [[ $? -ne 0 ]]; then
	 exit $?
fi

sudo docker build -t zyfdedh/fanach-dctl .

# docker run --rm -p 18080:8080 -v /var/run:/var/run zyfdedh/fanach-dctl
