#!/bin/sh

set -e

rm -rf ~/go/bin/minimalcli
go build main.go
cp main ~/go/bin/minimalcli