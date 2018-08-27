### centos7 编译
安装依赖
```shell
yum -y install epel-release && yum -y update
yum -y install gflags-devel snappy-devel zlib-devel bzip2-devel gcc-c++  libstdc++-devel
```
编译
```shell
export GOPATH=/your/go/path/directory  #设置GOPATH路径
cd $GOPATH/src
git clone https://github.com/okblockchainlab/go-nebulas.git ./github.com/nebulasio/go-nebulas
cd ./github.com/nebulasio/go-nebulas
export COIN_DEPS=`pwd`/depslib
./build.sh #run this script only if you first time build the project
./runbuild.sh
ls *.so
ls *.dylib
```


### 其它注意项
- createrawtransaction 每次生成的tx的timestamp不一样，所以会导致同样的输入，输出会有细微的差别
