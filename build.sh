#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o makespace .
echo "make success"
docker build -t makespace .
echo "build docker image success"
docker-compose up -d