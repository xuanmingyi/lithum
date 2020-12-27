#!/bin/bash

if [ -L $0 ]
then
    BASE_DIR=`dirname $(readlink $0)`
else
    BASE_DIR=`dirname $0`
fi


function build() {
    cd $BASE_DIR/../frontend
    npm run build:prod
    tar vzcf frontend.tar.gz dist/*
}

build
