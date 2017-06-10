@echo off
set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
go test -v ./...