@echo off
set GOOS=windows
set GOARCH=386
REM Event Build
go build -o FRCEventKiosk.exe -ldflags="-H windowsgui -X radioconfigtool.EventMode=true"
REM Home Use Build
go build -o FRCHomeKiosk.exe -ldflags="-H windowsgui -X radioconfigtool.EventMode=false"
