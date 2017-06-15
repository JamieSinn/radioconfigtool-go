@echo off
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

REM Build Resource File
echo Building Embedded Resource File (Manifest and icon)
_buildtools\rsrc_windows_386.exe -manifest="_buildtools\radioconfigtool.manifest" -ico="_buildtools\frc_64x64_1Fb_icon.ico" -o="rsrc.syso"

REM Debug Builds
echo Building Event Debug Build...
go build -o _dist/FRCEventDebug.exe -ldflags="-X firstinspires.org/radioconfigtool/config.eventmode=true"
echo Building Home Debug Build...
go build -o _dist/FRCHomeDebug.exe -ldflags="-X firstinspires.org/radioconfigtool/config.eventmode=false -X firstinspires.org/radioconfigtool/config.ENCRYPTION_KEY=IzLNm4rZK77TBCXopuRhufEP7x6UBOWl"

REM Event Build
echo Building Event Kiosk...
go build -o _dist/FRCEventKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=true -s -w" -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH

REM Home Use Build
echo Building Home Kiosk...
go build -o _dist/FRCHomeKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=false -X firstinspires.org/radioconfigtool/config.ENCRYPTION_KEY=IzLNm4rZK77TBCXopuRhufEP7x6UBOWl -s -w" -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH
echo Done.