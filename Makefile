
go-install:
	sh install-go.sh

all: go-install service test callApi

up: service test callApi

service:
	sh ./service.sh
test:
	sh ./test.sh
callApi:
	sh ./callApi.sh
