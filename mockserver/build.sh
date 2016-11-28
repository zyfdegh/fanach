#!/bin/sh

go build -o bin/mockserver

sudo docker build -t zyfdedh/fanach-mockserver .
