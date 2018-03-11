## 二进制协议

如tcp协议。tcp frame.

```shell
tcpdump -i any -nn -A port 7070

```

## [文本协议与二进制协议的选择](http://www.cnblogs.com/houkui/p/4268233.html)

http://www.vants.org/?post=106

http://blog.csdn.net/thinkry/article/details/41345881



#### 主程的晋升系统

http://blog.csdn.net/thinkry/article/details/41039679
主程的晋升攻略(2)：技术篇概要
主程的晋升攻略(3)：IP、DNS和CDN
主程的晋升攻略(4)：TCP、消息分包和协议设计
主程的晋升攻略(5)：HTTP协议和二进制协议的对比
主程的晋升攻略(6)：CGI和FastCGI



path 用于处理"/"分隔的串。因此可以处理http相关的路径。主要有2个: filepath.Dir(), filepath.Base()。

path/filepath用于处理操作系统相关的路径，可以跨平台。

code: homework/path.go



### 学生系统，函数的高级用法

refer: student.go

golang的面向对象是通过延迟绑定对象来实现的。

结构体实现了Update(int)方法。其实质却是一个func(s *Student, id int)函数。

```go
// 方法一
s := Student{Name: "binggan"}
f = s.Update 	// 即时绑定。即f已经确定的绑定到s这个实例上了

// 方法二
var f1 func(s *Student, id int)	// f1是一个函数，参数一是结构
f1 = (*Student).Update	// 	将f2绑定到结构体Student，而不是其实例！！！
f1(&s, 5)		// 调用f1时，才绑定到s这个实际实例，即延时绑定。

// 另定义一个实例
s1 := Student{Name: "jack"}
f1(&s1, 4)	// 将f1绑定到另一个实例上。
```

