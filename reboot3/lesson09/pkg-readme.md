## 包管理

相关变量

```shell
echo $GOPATH
~/jungleGOCode
# 工作路径，也是import包时的搜索路径。也是编译后的二进制存在顶级目录

echo $GOROOT
~/go1.10
# golang的安装路径，提供相关二进制命令
```

GOPATH项目管理目录。编译相关的目录。

- bin：存在编译后的可执行文件
- pkg：编译后的库文件
- src：存储golang的库文件



### pkgbase包

pkgbase/a/a.go

pkgbase/b/b.go

main1.go

* 引用上述包。正常引用，
* 按顺序调用每个被引用包的init()

```shell
# go run main1.go
init a
init b
a
b
```

main2.go

* 在main2中增加init，该函数被最后调用
* 使用`.`引用，被引用包的函数和变量做为主函数的一部分。

```shell
# go run main2.go
init a
init b
I'm main()
a
b
```

main3.go

* 使用`_`引用，只执行被引用包的init，不能调用其他函数。

```shell
# go run main3.go
init a
init b
I'm init in main
in main()
```



包的全局变量

* pkgbase/c/c.go中使用了全局变量

```shell
# go run main4.go
init a
global var in c
I'm init in main
in main()
c
```



**包引用的几种办法**

- 直接引用
- 引用时前面加字符串进行改名，避免namespace冲突
- 引用时，前面加`.`，将包中的变量和函数做为main的一部分直接使用
- 引用时，前面加`_`，只引用包，并做初始化
- 先引用先初始init()。链式引用
- init的顺序，全局变量的引用。init中可使用全局变量
- main()中也加入init()




### 包的可见性

引用“pkgbase/meta”包。其中一个变量可见，一个不可见。

```shell
# go run access1.go
&{0 100}
&{5 100}
&{6 0}
&{16 19}
```



internal目录

```shell
# go run access2.go
access2.go:4:2: use of internal package not allowed
```

包的依赖路径上包括了internal，就不能调用。



```shell
# go run access3.go
A
```

依赖的是"./internal/a"。引用正常。

也就是，除了相对路径外，其他方式不能包括`internal`字样。



#### 外部包

先实现一个独立的包。github.com/jkak/test/mytest。

继续在pkgtest目录操作：

```shell
# go get -v github.com/jkak/test/

# go run main5.go
hello golang!
```



如何保留当前使用的版本呢？

golang提供了vendor目录。社区也提供了多种管理工具。主要有glide, godep, dep等工具。

参考：
$GOPATH/src/github.com/jkak/learn-go/reboot3/pkgMgr

