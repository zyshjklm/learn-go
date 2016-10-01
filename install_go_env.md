
# install for mac

##  download 

http://www.golangtc.com/download

## env

```bash

mkdir $HOME/_go/{bin,src,pkg}

# for iTerm2 and  Zsh
vi ~/.zshrc
# add below

export GOPATH=$HOME/_go
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin


source ~/.zshrc

```

go env 

```shell
#
# GOPATH="/Users/song/_go"
# GORACE=""
# GOROOT="/usr/local/go"
# 
```



## test go

```bash

vi ~/test.go

# add below
package main

import "fmt"

func main() {
	fmt.Printf("Hello, World!\n")
}
### end 

go run ~/test.go
```

## install gosublime

refer: http://my.oschina.net/Obahua/blog/110767
step 7.


## vim

refer:
https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/01.4.md

其中，1.6.2版本的go，直接下载pkg包安装后，是没有$GOROOT/misc/vim/目录的。
参考：http://www.tuicool.com/articles/nmqaMbq
下载1.3.3版本的源文件。解压后即可得到。



## Windows7 install go1.7

* 安装：
  * 下载 go1.7.windows-amd64.msi。
  * 双击安装，默认安装在C:\Go\目录下。并自动设置了环境变量：
    * GOROOT= "C:\Go\"
    * PATH变量中添加了"C:\Go\bin\"



```shell
# open cmd terminal
go 

go env
# check :
# GOPATH="", GOROOT=C:\Go\
```



* 设置环境变量：
  * 创建编译用的工作目录。比如创建为c:\work\目录。
  * 设置其它环境变量，进入环境变量设置页面，添加：GOPATH=C:\work\

```shell
go env 
# check GOPATH="C:\work\"
```



环境变量入口，有两方式：

* 在“计算机”上右键，进入“属性”，-> “高级系统设置” -> “高级” -> 环境变量
* 控制面板，系统和安全，系统，高级系统设置

调试程序如前面hello.go。





