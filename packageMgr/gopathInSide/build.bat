@echo off
set OLD_GO_PATH=%GOPATH%
set OLD_GO_BIN=%GOBIN%

set CURRENT_PATH=%CD%
set SRC_PATH=%CURRENT_PATH%\src
set APP_NAME=myapp
set GOPATH=%CURRENT_PATH%
set GOBIN=%GOPATH%\bin\
mkdir %GOBIN%
echo "must using this to write[go env] in windows"
go env -w GOBIN=%GOBIN%
go env


echo formating...
gofmt -w %SRC_PATH%

echo building...
go clean
go install "%APP_NAME%"
echo build done...


set GOPATH="%OLD_GO_PATH%"
set GOBIN="%OLD_GO_BIN%"
