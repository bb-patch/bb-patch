#!/usr/bin/env bash

echo "Building application"
go build -ldflags "-s -w -H=windowsgui" -o bb-patch-1.00.exe