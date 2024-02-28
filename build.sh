#!/bin/bash

# 【linux/amd64】
echo "start build linux/amd64 ..."
GOOS=linux GOARCH=amd64 go build -o pls_linux_amd64 .

# 【darwin/amd64】
echo "start build darwin/amd64 ..."
GOOS=darwin GOARCH=amd64 go build -o pls_darwin_amd64 .

# 【windows/amd64】
echo "start build windows/amd64 ..."
GOOS=windows GOARCH=amd64 go build -o pls_windows_amd64.exe .

echo "all binary files are listed below"
ls | grep pls
