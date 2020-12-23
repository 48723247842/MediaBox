#!/bin/bash
# rm go.mod
# rm go.sum

# sudo rm -rf ~/go/pkg/mod/github.com/0187773933/ || echo ""
# sudo rm -rf  /home/morphs/go/src/$1 || echo ""
# sudo rm -rf  /home/morphs/go/pkg/mod/$1 || echo ""
# go clean --modcache

# go mod init c2server

# We have to force golang to not care about cache of repos we are currently updating
RMUHash=$(curl -s 'https://api.github.com/repos/0187773933/RedisManagerUtils/git/refs/heads/master' | jq -r '.object.sha')
go get "github.com/0187773933/RedisManagerUtils/@$RMUHash"

VLCHash=$(curl -s 'https://api.github.com/repos/0187773933/VLCWrapper/git/refs/heads/master' | jq -r '.object.sha')
go get "github.com/0187773933/VLCWrapper/@$VLCHash"

VizioHash=$(curl -s 'https://api.github.com/repos/0187773933/VizioController/git/refs/heads/master' | jq -r '.object.sha')
# go get "github.com/0187773933/VLCWrapper/@$VizioHash"
# go get "github.com/0187773933/VLCWrapper/@f1f7967ab6eff94c3770dad40ab77682fa081edf"
go get "github.com/0187773933/VLCWrapper@$VizioHash"

#/usr/local/bin/goBuildAllPlatforms mediabox
go run main.go