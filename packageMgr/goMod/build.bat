
@echo off
set CURRENT_PATH=%CD%
set GOBIN=%CURRENT_PATH%\bin
set SRC_PATH=%CURRENT_PATH%
mkdir %GOBIN%

echo formating...
gofmt -w %SRC_PATH%


go env -w GO111MODULE=on
go env -w GOBIN=%GOBIN%
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go clean
go env
go install
