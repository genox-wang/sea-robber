#!/bin/bash
export GOPATH=$(pwd)/sea-robber-api

echo "GOPATH=$GOPATH"

cd sea-robber-api/src/sea-robber-api

echo "copy glide mirror..."

cp ./.glide/mirrors.yaml ~/.glide/mirrors.yaml

echo "glide install..."

glide install

echo "build..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sea-robber-api ./app

chmod +x sea-robber-api

echo "build complete"