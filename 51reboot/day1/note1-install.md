
## golang 介绍



* 基础知识
* 高阶知识
  * Goroutine, channel
  * Lock, race, debug
  * Backage, test, benchmark
  * Reflect, cgo, unsafe
* 项目
  * 监控系统
  * agent
  * kafka传输
  * es存储



### go版本

go从1.0开始的代码，向下兼容。用最新的1.8.3可能编译1.0时的程序。

1.8.3.

第二位：8是每半年周期的发布版本。周年2次更新。

第三位：是修复版本。



### 安装

#### 编译安装

先安装一个早期版本，用于编译。

1.4.3是纯C写的。

```shell
# download
tar xf 
mv go go1.4.3
cd go1.4.3
cd src
./make.bash
# 进行编译。
bin/go version


```

下载最新的源码。1.8.3

```shell
curl https://dl.gocn.io/golang/1.8.3/go1.8.3.src.tar.gz
tar xf 
mv go go1.8
cd go1.8
cd src
./make.bash
# 提示要求设置一个环境变量，并要求编译环境要大于1.4
# $GOROOT_BOOTSTRAP

export GOROOT_BOOTSTRAP=$HOME/...
./make.bash

```

环境变量部分见下面。



#### 直接安装mac

```shell
wget https://dl.gocn.io/golang/1.8.3/go1.8.3.darwin-amd64.tar.gz
cd ~/Downloads
tar -zxvf go1.8.3.darwin-amd64.tar.gz
mv go ~/go1.8.3

go1.8.3/bin/go version
go version go1.8.3 darwin/amd64
chmod 775 go1.8.3

```

设置环境变量

```shell
vim ~/.zshrc

export GOROOT=/Users/jungle/go1.8.3/
export GOPATH=$HOME/jungleGo
export GOBIN=$GOPATH/bin

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

source ~/.zshrc

which go
/Users/jungle/go1.8.3//bin/go

go version
go version go1.8.3 darwin/amd64
```



旧的安装方法

```shell
from 1.7.4 to 1.8

mv /usr/local/go /usr/local/go1.7.4
vi ~/.zshrc
modify GOROOT
source ~/.zshrc

download http://www.golangtc.com/static/go/1.8/go1.8.darwin-amd64.tar.gz

mv go1.8.darwin-amd64.tar.gz /usr/local
cd /usr/local
tar -zxvf go1.8.darwin-amd64.tar.gz
mv go go1.8.1

cd go1.8.1/src
bash all.bash
```



### hello world

```go
package main

import "fmt"

func main () {
	fmt.Println("Hello golang")
}
	
```

编译运行

```shell
go run hello.go
Hello golang
```

