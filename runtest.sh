#!/bin/sh

if [ -z "$COIN_DEPS" ]; then
        printf "No COIN_DEPS detected!\\n"
        printf "Setup COIN_DEPS before build: export COIN_DEPS=`pwd`/depslib \\n"
        exit 1
fi

export CGO_CFLAGS="-I$COIN_DEPS/rocksdb/include -I$COIN_DEPS/snappy/include"
export CGO_LDFLAGS="-L$COIN_DEPS/rocksdb/lib -L$COIN_DEPS/snappy/lib"

go test -v ./okwallet/okwallet
