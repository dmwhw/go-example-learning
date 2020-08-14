@echo off
set GOBIN=%CD%\bin
mkdir %GOBIN%

go env -w GO111MODULE=on
go env -w GOBIN=%GOBIN%
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/