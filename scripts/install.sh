#!/bin/bash

if [ -L $0 ]
then
    BASE_DIR=`dirname $(readlink $0)`
else
    BASE_DIR=`dirname $0`
fi

function install () {
    cd $BASE_DIR/../frontend
    npm install --registry=https://registry.npm.taobao.org
}

install
