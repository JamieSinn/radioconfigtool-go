@echo off
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

build_debug.bat

REM Event Build
echo Building Event Kiosk...
go build -o _dist/FRCEventKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=true -s -w" -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH

REM Home Use Build
echo Building Home Kiosk...
go build -o _dist/FRCHomeKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=false -X firstinspires.org/radioconfigtool/config.ENCRYPTION_KEY=IzLNm4rZK77TBCXopuRhufEP7x6UBOWl -s -w" -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH
echo Done.