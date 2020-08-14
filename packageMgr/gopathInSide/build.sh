###验证通过
echo off
export OLD_GO_PATH=$GOPATH
export OLD_GO_BIN=$GOBIN

export CURRENT_PATH=`pwd`

export SRC_PATH=$CURRENT_PATH/src

export APP_NAME=myapp
export GOPATH=$CURRENT_PATH
export GOBIN=$GOPATH/bin
mkdir -p  $GOBIN

 
go env
echo formating...
gofmt -w $SRC_PATH

echo building...
go install $APP_NAME
echo build done...


export GOPATH=$OLD_GO_PATH
export GOBIN=$OLD_GO_BIN

#go get github.com/garyburd/redigo/redis
#export GOPROXY=https://mirrors.aliyun.com/goproxy/