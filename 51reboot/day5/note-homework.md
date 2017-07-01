
### homework first

学生管理系统，load之后，原有的数据是还在的？可能存在覆盖。还是需要先清空原有数据？


关于字符串的结构体

```go
type StringStruct struct {
   pointer *[]string
   len uint64
   cap uint64
}
```





如何去掉行尾的换行？

1 如果确切知道行尾必然有换行，则直接切片。
s := "hello\n"
s = s[:len(s)-1]

2 string.StrimSpace()

3 strings.Replace(line, "\n", "", 1)