@echo off
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

REM Build Resource File
echo Building Embedded Resource File (Manifest and icon)
_buildtools\rsrc_windows_386.exe -manifest="_buildtools\radioconfigtool.manifest" -ico="_buildtools\frc_64x64_1Fb_icon.ico" -o="rsrc.syso"

REM Event Build
echo Building Event Kiosk...
go build -o _dist/FRCEventKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=true -s -w"
REM Home Use Build
echo Building Home Kiosk...
go build -o _dist/FRCHomeKiosk.exe -ldflags="-H windowsgui -X firstinspires.org/radioconfigtool/config.eventmode=false -s -w"
echo Done.