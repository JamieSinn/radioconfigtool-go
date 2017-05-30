@echo off
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
REM Event Build
go build -o dist/FRCEventKiosk.exe -ldflags="-H windowsgui -X config/config.EventMode=true"
REM Home Use Build
go build -o dist/FRCHomeKiosk.exe -ldflags="-H windowsgui -X config/config.EventMode=false"
