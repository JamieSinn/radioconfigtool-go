@echo off
set GOOS=windows
set GOARCH=386
go build -ldflags="-H windowsgui"
