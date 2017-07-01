



Return code

```shell

go run return.go
open abcd.txt: no such file or directory
exit status 2

go build return.go

./return
open abcd.txt: no such file or directory

echo $?
2
```



## 闭包

函数及其上下文。

一个函数被返回后，带有上下文。该函数可以使用在其函数以外定义的相关环境和参数。

闭包的坑：

