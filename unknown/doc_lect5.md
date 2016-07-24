### lect5 goto if for

### switch 

```go
// switch
// 不需要写break，会自动终止。
// 如果希望继续下一个case,需要使用fallthrough
	
```

### goto break continue

```go
/*
goto, break, continue
三者都 可以配合标签使用
标签区分大小写。
break, continue可以配合标签用于多层循环的跳出。

但go却只是改变执行的位置。

用break, continue，标签是位置在上面，
用goto, 标签是在位置的后面。
*/
```



### pointer

```go
/*
指针：
虽然保留了，但不支持 指针运算以及" -> "
取成员变量使用”.“
 
默认值是nil

另外，++， --是语句，而不是表达式。
表达式可以放在 = 的右边。因此：
	a++可以做为单独的一行，但不能使用b = a++这样。
*/

```

