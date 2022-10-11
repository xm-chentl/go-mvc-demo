#!/bin/bash

cur_dir=$(pwd)
function build(){
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ${1}/${1}.proto
    mkdir -p go/$1
    mv ${1}/*.go $cur_dir/go/$1
}

function build_all(){
    dirs=$(ls -l |awk '/^d/ {print $NF}' | grep -ve '^go$')
    for d in $dirs
    do
        build $d
    done
}
if [ "$1" == "" ];then
    echo 请传入指定参数
    echo 如: sh build.sh 指定目录
    echo all 默认根目录下所有目录、指定目录 如 common
elif [ "$1" == "all" ];then
    build_all
else
    for a in $*; do
        build $a
    done
fi

    