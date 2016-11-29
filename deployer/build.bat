@echo off

go get golang.org/x/crypto/ssh
go build -o bin/deployer.exe
cp -r static bin/
