#!/bin/bash

GOOS=darwin GOARCH=arm64 go build -o ../bin/cringecast-server main.go
GOOS=windows GOARCH=amd64 go build -o ../bin/cringecast-server-win.exe main.go