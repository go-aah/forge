#!/usr/bin/env bash

set -e
echo "Preparing aah framework environment for travis"

# export variables
export PATH=$GOPATH/bin:$PATH
export CURRENT_BRANCH=$TRAVIS_BRANCH
export AAH_SRC_PATH=$GOPATH/src/aahframework.org
export AAH_PKG_NAME=$1
export AAH_PKG_VERSION=$2
export AAH_PKG=$AAH_PKG_NAME.$AAH_PKG_VERSION

# create aah src path
mkdir -p $AAH_SRC_PATH

# preparing go sources
src_module_path=$GOPATH/src/github.com/go-aah/$AAH_PKG_NAME
target_module_path=$AAH_SRC_PATH/$AAH_PKG

mkdir -p "$target_module_path"
cp -a "$src_module_path/." "$target_module_path/"

# update travis build directory
export TRAVIS_BUILD_DIR=$target_module_path
echo "TRAVIS_BUILD_DIR: $TRAVIS_BUILD_DIR"

# cleanup and change directory
rm -rf $GOPATH/src/github.com/go-aah/*
