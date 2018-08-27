#!/bin/sh

PWD=`pwd`
LIBDOWNLOAD=$PWD/libdownload

if [ -z "$COIN_DEPS" ]; then
        printf "No COIN_DEPS detected!\\n"
        printf "Setup COIN_DEPS before build: export COIN_DEPS=`pwd`/depslib \\n"
        exit 1
fi

if [ ! -d "$LIBDOWNLOAD" ];then
        mkdir $LIBDOWNLOAD
fi

if [ ! -d "$COIN_DEPS" ];then
        mkdir $COIN_DEPS
fi

cd $COIN_DEPS

if [ ! -d ./rocksdb ];then
        mkdir rocksdb
fi

if [ ! -d ./snappy ];then
        mkdir snappy
fi

if [ ! -d ./bzip2 ];then
        mkdir bzip2 
fi


cd $LIBDOWNLOAD

if [ ! -f ./rocksdb-5.14.3.tar.gz ];then
        wget https://github.com/facebook/rocksdb/archive/rocksdb-5.14.3.tar.gz
        if [ ! -f ./rocksdb-5.14.3.tar.gz ];then
                echo "Error cannot download rocksdb-5.14.3.tar.gz" >&2
                exit 1
        fi
fi

if [ ! -d ./rocksdb-rocksdb-5.14.3 ];then
        tar xf rocksdb-5.14.3.tar.gz
        if [ ! -d rocksdb-rocksdb-5.14.3 ];then
                echo "Error cannot tar xf rocksdb-5.14.3.tar.gz"
                exit 1
        fi
fi

cd rocksdb-rocksdb-5.14.3
CFLAGS="-fno-strict-aliasing -fPIC" make static_lib
INSTALL_PATH=$COIN_DEPS/rocksdb make install-static


cd $LIBDOWNLOAD

if [ ! -f ./snappy-1.1.7.tar.gz ];then
        wget https://github.com/google/snappy/archive/1.1.7.tar.gz -O snappy-1.1.7.tar.gz
        if [ ! -f ./snappy-1.1.7.tar.gz ];then
                echo "Error cannot download snappy-1.1.7.tar.gz" >&2
                exit 1
        fi
fi

if [ ! -d ./snappy-1.1.7 ];then
        tar xf snappy-1.1.7.tar.gz
        if [ ! -d snappy-1.1.7 ];then
                echo "Error cannot tar xf snappy-1.1.7.tar.gz"
                exit 1
        fi
fi

cd snappy-1.1.7
mkdir build
cd build
cmake -DCMAKE_INSTALL_PREFIX=$COIN_DEPS/snappy ..
make
make install


cd $LIBDOWNLOAD

if [ ! -f ./bzip2-1.0.2.tar.gz ];then
        wget ftp://sources.redhat.com/pub/bzip2/v102/bzip2-1.0.2.tar.gz
        if [ ! -f ./bzip2-1.0.2.tar.gz ];then
                echo "Error cannot download bzip2-1.0.2.tar.gz" >&2
                exit 1
        fi
fi

if [ ! -d ./bzip2-1.0.2 ];then
        tar xf bzip2-1.0.2.tar.gz
        if [ ! -d bzip2-1.0.2 ];then
                echo "Error cannot tar xf bzip2-1.0.2.tar.gz"
                exit 1
        fi
fi

cd bzip2-1.0.2
make libbz2.a && make install PREFIX=$COIN_DEPS/bzip2


cd $LIBDOWNLOAD/..

if [ ! -f ./vendor.tar.gz ];then
        wget http://ory7cn4fx.bkt.clouddn.com/vendor.tar.gz
        if [ ! -f ./vendor.tar.gz ];then
                echo "Error cannot download vendor.tar.gz" >&2
                exit 1
        fi
fi

if [ ! -d ./vendor ];then
        tar xf vendor.tar.gz
        if [ ! -d vendor ];then
                echo "Error cannot tar xf vendor.tar.gz"
                exit 1
        fi
fi
