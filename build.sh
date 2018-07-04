#!/usr/bin/env bash
#!/bin/bash
export GOPATH=$GOPATH:`pwd`
export GOARCH=amd64

go get -u github.com/armon/go-socks5

export GOOS=linux
go build -ldflags "-w -s -linkmode internal" -o sproxy-linux src/sproxy.go
export GOOS=darwin
go build -ldflags "-w -s -linkmode internal" -o sproxy-mac src/sproxy.go
export GOOS=windows
go build -ldflags "-w -s -linkmode internal" -o sproxy-win.exe src/sproxy.go
