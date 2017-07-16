## go里面的unix的哲学

unix一切皆文件。go继承了这一哲学。下面演示关于文件相关的操作。

流式的处理。



#### 1 遍历文件

pathWalk.go 接受一个目录参数，然后遍历显示所有的目录。



#### 2 压缩与解压

通过tar czf登陆压缩包，再用程序来解压。解压时使用了gzip.NewReader(os.Stdin)。

从标准输入读取压缩文件，经过gzip.NewReader()接口后，即得到了解压后的数据。

zcat.go实现了该功能。类似linux的zcat命令。

```shell
# tar cf zcat.tar zcat.go
# file zcat.tar
zcat.tar: POSIX tar archive

# tar czf zcat.tar.gz zcat.go
# file zcat.tar.gz
zcat.tar.gz: gzip compressed data, last modified: Wed Jul 12 23:04:12 2017, from Unix

# gzip -l zcat.tar.gz
  compressed uncompressed  ratio uncompressed_name
         264        10240  97.4% zcat.tar
# ls -l zcat.*
-rw-r--r--  1 song  staff   169 Jul  8 18:04 zcat.go
-rw-r--r--  1 song  staff  2048 Jul 13 07:03 zcat.tar
-rw-r--r--  1 song  staff   264 Jul 13 07:04 zcat.tar.gz

# go build zcat.go

# ./zcat < zcat.tar
2017/07/13 07:06:40 gzip: invalid header

# ./zcat < zcat.tar.gz
zcat.go000644 000765 000024 00000000251 13130127050 012445 0ustar00songstaff000000 000000 package main

import (
### ...
	r := gzip.NewReader(os.Stdin)
### ...	

```



#### 2 打包，压缩

tar即Tape archive。

打包像磁带一样，是顺序存储的。存储的格式就是：

* 文件头，文件内容；
* 文件头，文件内容；
* 文件头，文件内容；。。。

archive/tar的tar.Next()用来读取文件头。返回文件头Header.

根据指针读取文件，直到EOF。

* **tar-listFile.go** 列取tar包中的文件及目录名。



```shell
# tar cf interface.tar ../method-pointer.go
# tar cf interface.tar ../interface

# go run tar-listFile.go < interface.tar
../method-pointer.go

# go run tar-listFile.go < interface.tar
../interface/
../interface/blankInferface.go
../interface/ByteCounter-interface-4.go
../interface/ByteCounter-struct-2.go
#### ...
```

* tar-unTar.go 解tar包

```shell
pwd
learn-go/51reboot/day6/allAreFiles
cd ..
tar cf interface.tar interface
find interface -type f | xargs md5

mv interface.tar allAreFiles
cd allAreFiles
# 解压并验证md5，确认与原始目录下的md5相同。
go run tar-unTar.go < interface.tar
find interface -type f | xargs md5

```





写压缩，tar.Writer...

解压只是在压缩的前面加了一个gzip.NewReader() 进行一轮过滤。把得到的接口类型再传给后面的实现。

```go
    //untar.go

	uncompress, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(uncompress)

    // 解压：src.tar.gz -> gzip.NewReader()  -> tar.NewReader() -> src (*os.File)

    // 压缩：src1 src2. -> tar.NewReader()  -> gzip.NewReader() -> fd.tar.gz (*os.File)
```

### 装饰器对象

MyReader()

函数不能记录状态，只能使用全局变量。而使用装饰器对象，则可以记录。

比如实现一个自己的Reader，对读过的字节计数。

#### 1）myReader-base.go

```shell
### prepare for test
cp zcat.go test-zcat.go
md5 zcat.go test-zcat.go
#MD5 (zcat.go) = 529eae23d8b3662e8bddb75a0a5583e6
#MD5 (test-zcat.go) = 529eae23d8b3662e8bddb75a0a5583e6
tar -czf kkk.tar.gz test-zcat.go
ls -l kkk.tar.gz
#-rw-r--r--  1 song  staff  297 Jul 16 14:28 kkk.tar.gz
gzip -l kkk.tar.gz
#  compressed uncompressed  ratio uncompressed_name
#         297        10240  97.1% kkk.tar
file kkk.tar.gz
#kkk.tar.gz: gzip compressed data, last modified: Sun Jul 16 06:28:53 2017, from Unix
rm test-zcat.go

go run myReader-base.go < kkk.tar.gz
#test-zcat.go000644 000765 000024 00000000323 13132603617 013434 0ustar00songstaff000000 000000 package main
# ...
# copied num: 10240; rd size: 10240
```

如上可见，zcat.go --> test-zcat.go -> kkk.tar -> kkk.tar.gz

其中从gzip命令可以看到kkk.tar的原始大小是10240字节。这和myReader-base.go的结果相同。这个脚本是直接读的gzip解压后的包。没有处理包关。

该文件中，Reader的功能和被注释掉的一行// io.Copy(os.Stdout, uncompress)，除了统计外，基本功能是一样的。

#### 2）myReader-adv.go

```shell
go run myReader-adv.go < kkk.tar.gz
# ...

# copied num: 211; rd size: 211

ls -l test-zcat.go
-rw-r--r--  1 song  staff  211 Jul 16 15:53 test-zcat.go

md5 zcat.go test-zcat.go
MD5 (zcat.go) = 529eae23d8b3662e8bddb75a0a5583e6
MD5 (test-zcat.go) = 529eae23d8b3662e8bddb75a0a5583e6
```

第一个文件是直接将tar的Reader给了自定义Reader，而第二个则是先处理了tar的包头，即先tar.Next()之后，剩下的才是包体。从而计数读到的大小和文件大小相同。

