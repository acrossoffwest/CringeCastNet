#!/bin/bash

GOOS=darwin GOARCH=arm64 go build -o ../bin/cringecast-client main.go
GOOS=windows GOARCH=amd64 go build -o ../bin/cringecast-client-win.exe main.go