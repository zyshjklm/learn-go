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

函数不能记录状态，只能使用全局变量。


