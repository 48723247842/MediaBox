#!/bin/bash
rm go.mod
rm go.sum

sudo rm -rf ~/go/pkg/mod/github.com/0187773933/ || echo ""
sudo rm -rf  /home/morphs/go/src/$1 || echo ""
sudo rm -rf  /home/morphs/go/pkg/mod/$1 || echo ""
go clean --modcache

go mod init universalstatuspoller

# We have to force golang to not care about cache of repos we are currently updating
RMUHash=$(curl -s 'https://api.github.com/repos/0187773933/RedisManagerUtils/git/refs/heads/master' | jq -r '.object.sha')
go get "github.com/0187773933/RedisManagerUtils/@$RMUHash"


BinaryName="mediaboxUniversalStatusPoller"
rm -rf ./bin/
# https://stackoverflow.com/questions/25051623/golang-compile-for-all-platforms-in-windows-7-32-bit
# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

declare -a linux_architectures=(
	"amd64"
	"arm"
	"arm64"
)
declare -a darwin_architectures=(
	"amd64"
)
declare -a windows_architectures=(
	"amd64"
)

for architecture in "${linux_architectures[@]}"
do
	echo "Building Linux: $architecture"
	GOOS=linux GOARCH=$architecture go build -o bin/linux/$architecture/$BinaryName
done

for architecture in "${darwin_architectures[@]}"
do
	echo "Building Darwin: $architecture"
	GOOS=darwin GOARCH=$architecture go build -o bin/darwin/$architecture/$BinaryName
done

for architecture in "${windows_architectures[@]}"
do
	echo "Building Windows: $architecture"
	GOOS=windows GOARCH=$architecture go build -o bin/windows/$architecture/$BinaryName
done

sudo cp ./bin/linux/amd64/mediaboxUniversalStatusPoller /usr/local/bin/