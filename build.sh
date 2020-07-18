#!/usr/bin/env sh

go mod download
go build -ldflags "-s -w" -o build/