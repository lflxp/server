#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
docker stop servertest
docker rm servertest
docker build -t lxp/servertest .
docker run -d --name -P servertest lxp/servertest
