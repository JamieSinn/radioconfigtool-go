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