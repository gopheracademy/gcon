#!/bin/bash

set -e
export GOPATH=~/go
cd ~/go/src/github.com/gopheracademy/gcon
git pull
/usr/local/go/bin/go build
export GO_ENV=production
/root/go/bin/buffalo soda migrate up
systemctl restart gcon
