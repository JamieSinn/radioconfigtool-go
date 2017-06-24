@echo off
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1

REM Go to the test client directory
cd _testclient
REM Building Programmer test client
echo Building programmer...
cd programmer
go build -o ..\..\_dist/programmer_test.exe

cd ..

REM Building TFTP test client
echo Building tftp...
cd tftp
go build -o ..\..\_dist/om5pac_tftp_test.exe

cd ..\..

