#!/bin/bash

# build SaveHelloFunction
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o lib/lambda/main lib/lambda/main.go 