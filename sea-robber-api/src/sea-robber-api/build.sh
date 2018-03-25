#!/bin/bash
export GOPATH=$(pwd)/../../

echo "GOPATH=$GOPATH"

echo "copy glide mirror..."

cp ./.glide/mirrors.yaml ~/.glide/mirrors.yaml

echo "glide install..."

glide install

echo "build..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sea-robber-api ./app

chmod +x sea-robber-api

echo "build complete"