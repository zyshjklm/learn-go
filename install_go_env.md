
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



