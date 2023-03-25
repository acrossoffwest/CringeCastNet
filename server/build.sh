#!/bin/bash

GOOS=darwin GOARCH=arm64 go build -o ../bin/cringecast-server main.go
GOOS=windows GOARCH=amd64 go build -o ../bin/cringecast-server-win.exe main.go
GOOS=linux GOARCH=amd64  go build -o ../bin/cringecast-server-linux main.go
GOOS=linux GOARCH=arm64  go build -o ../bin/cringecast-server-linux-arm main.go