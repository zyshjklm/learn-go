

### fmt

- Print
- Printf
- Println

另有前缀：F, S。

Fprint系统是写到文件句柄。而Sprint系统是格式化成一个串。

注意差别

```go
fmt.Println(fmt.Sprint("http://%s/%s", "www.baidu.com", "about.html"))
fmt.Println("http://%s/%s", "www.baidu.com", "about.html")
// 前2个效果相同。下面一个才是格式化字符串。
fmt.Println(fmt.Sprintf("http://%s/%s", "www.baidu.com", "about.html"))
```

Code: fmt/main.go



### shell

通过循环和读取输入，解析输入来实现shell。

有2个方面来实现读取：

- 通过最简单的fmt.Scan(&str, &n)
- 通过f = bufio.NewReader(os.Stdin); f.ReadString('\n')

第二种方式更能控制读取一行。