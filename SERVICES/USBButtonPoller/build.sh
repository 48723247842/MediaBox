#!/bin/bash
rm -rf ./bin/
go clean --modcache
GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/mediaboxUSBButtonPoller
sudo cp bin/linux/amd64/mediaboxUSBButtonPoller /usr/local/bin/mediaboxUSBButtonPoller