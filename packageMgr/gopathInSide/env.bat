@echo off
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

set OLD_GO_PATH=%GOPATH%
set OLD_GO_BIN=%GOBIN%

set CURRENT_PATH=%CD%
set SRC_PATH=%CURRENT_PATH%\src
set APP_NAME=myapp
set GOPATH=%CURRENT_PATH%
