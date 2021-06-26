#!/bin/bash

export PATH=$GOROOT/bin:$PATH

go get -u github.com/fvbock/endless
go get -u github.com/gin-gonic/gin
go get -u github.com/google/gops/agent
go get -u github.com/itsjamie/gin-cors
#go get -u xi2.org/x/logrot
go build main.go
sleep 1
nohup ./main &
sleep 1
echo "Trivago service is running on port 9000"
