#!/bin/sh

echo "Removing old binaries"
mkdir -p ../dist
rm ../dist/sleep-on-lan-* > /dev/null 2>&1

echo "Building Linux binaries"

GOOS=linux GOARCH=amd64 go build -o ../dist/sleep-on-lan-linux-amd64 ..
GOOS=linux GOARCH=arm go build -o ../dist/sleep-on-lan-linux-arm ..
GOOS=linux GOARCH=arm64 go build -o ../dist/sleep-on-lan-linux-arm64 ..

echo "Building Windows binary"

GOOS=windows GOARCH=amd64 go build -o ../dist/sleep-on-lan-windows-amd64.exe ..
