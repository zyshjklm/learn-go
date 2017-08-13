## 加密代理隧道

之前一共实现了两类代理。

* tcp层代理，可用于所有使用tcp协议的应用。http, ssh, ftp等
* 使用socks5代理。


下面讲对tcp的流式加密。并和socks5配合使用。



目录：51reboot/cryptoSocks5-V1/

### 1）中间件mycrypto.go

**结构体：**

* 封装一个Writer结构体，并提供rc4加密。
* 封装一个Reader结构体，并提供一个rc4加密。
* 并分别提供构造方法，以及对应的接口。

**测试**：

* 基准测试
* 性能测试

**代码：**

* ./mycrypto/mycrypto.go
* ./mycrypto/mycrypto_test.go



测试过程：

```shell
#### pwd : cryptoSocks5-V1

# cd mycrypto
# go test
PASS
ok  	github.com/jungle85gopy/learn-go/51reboot/cryptoSocks5-V1/mycrypto	0.009s

# go build
# echo "abcd" | ./mycrypto | ./mycrypto
abcd

#### 测试随机文件

# dd if=/dev/urandom of=./block.0 bs=1m count=20
20+0 records in
20+0 records out
20971520 bytes transferred in 1.967966 secs (10656444 bytes/sec)

# time ./mycrypto < block.0 > block.1
./mycrypto < block.0 > block.1  0.05s user 0.05s system 80% cpu 0.124 total 

# time ./mycrypto < block.1 > block.2
./mycrypto < block.1 > block.2  0.06s user 0.06s system 65% cpu 0.187 total

# ls -l block*
-rw-r--r--  1 song  staff  20971520 Aug 13 23:43 block.0
-rw-r--r--  1 song  staff  20971520 Aug 13 23:44 block.1
-rw-r--r--  1 song  staff  20971520 Aug 13 23:44 block.2

# md5 block*
MD5 (block.0) = 2f781e483d955a580eb407119972e348
MD5 (block.1) = 303d1985212c4eb20d06104e943d58b6
MD5 (block.2) = 2f781e483d955a580eb407119972e348
#### block.0和两次加密后的文件block.2相同。

#### 性能测试

# go test -bench .
go test -bench .
BenchmarkCrypto-4     500000	   2241 ns/op	 456.77 MB/s
PASS
ok  	github.com/jungle85gopy/learn-go/51reboot/cryptoSocks5-V1/mycrypto	1.156s
```

解释：

* 2241 ns/op单次时延，单位为纳秒
* 456.77 MB/s 每秒加密的流量




