yum -y install epel-release && yum -y update
yum -y install gflags-devel snappy-devel zlib-devel bzip2-devel gcc-c++  libstdc++-devel

export COIN_DEPS=`pwd`/depslib

SendRawTransaction

createrawtransaction 每次的timestamp不一样，所以会导致同样的输入，输出会有细微的差别
