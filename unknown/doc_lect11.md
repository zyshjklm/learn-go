## lect11 method



### 结构的方法

```go

type A struct {
	Name string
}

func (a A) Print() {
	fmt.Println("A")
}

func main() {
	a := A{}
	a.Print()
}
```

