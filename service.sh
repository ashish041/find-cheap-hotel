#!/bin/bash

export PATH=/usr/local/go/bin:$PATH
go mod init main
go mod tidy
go build main.go
sleep 1
nohup ./main &
sleep 1
echo "Trivago service is running on port 9000"
