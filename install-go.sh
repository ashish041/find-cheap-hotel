#!/bin/bash

echo "Installing Go...."
rm -f go1.16.linux-amd64.tar.gz
wget https://dl.google.com/go/go1.16.linux-amd64.tar.gz
sudo sh -c 'sudo tar -C /usr/local/ -xzf go1.16.linux-amd64.tar.gz'

echo GOROOT=/usr/local/go
echo PATH=/usr/local/go/bin:$PATH
echo GOBIN=${GOBIN:-$(pwd)}
echo export PATH GOROOT GOBIN

go version
echo "Done"
