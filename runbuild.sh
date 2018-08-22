#!/bin/sh

if [ -z "$COIN_DEPS" ]; then
  printf "No COIN_DEPS detected!\\n"
  printf "Setup COIN_DEPS before build: export COIN_DEPS=`pwd`/depslib \\n"
  exit 1
fi

if [ -z "$JAVA_HOME" ]; then
  printf "No JAVA_HOME detected! "
  printf "Setup JAVA_HOME before build: export JAVA_HOME=/path/to/java\\n"
  exit 1
fi

VERSION=1.0.8
COMMIT=`git rev-parse HEAD`
BRANCH=`git rev-parse --abbrev-ref HEAD`

LDFLAGS="-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.compileAt=`date +%s`"

EXT=so
NM_FLAGS=
TARGET_OS=`uname -s`
case "$TARGET_OS" in
  Darwin)
    EXT=dylib
    export CGO_CFLAGS="-I$COIN_DEPS/rocksdb/include -I$COIN_DEPS/snappy/include -I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin"
    ;;
  Linux)
    EXT=so
    NM_FLAGS=-D
    export CGO_CFLAGS="-I$COIN_DEPS/rocksdb/include -I$COIN_DEPS/snappy/include -I${JAVA_HOME}/include -I${JAVA_HOME}/include/linux"
    ;;
  *)
  echo "Unknown platform!" >&2
  exit 1
esac

export CGO_LDFLAGS="-L$COIN_DEPS/rocksdb/lib -L$COIN_DEPS/snappy/lib"


go build -o libneb.${EXT} -buildmode=c-shared -ldflags="${ldflags}" ./okwallet/libneb
[ $? -ne 0 ] && exit 1
nm ${NM_FLAGS} libneb.${EXT} |grep "[ _]Java_com_okcoin"
