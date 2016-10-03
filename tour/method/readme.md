## method

go does not have classes. you can define methods on struct type.

the method receiver appears in its own argument list between the func keyword and the method name.

you can declare a method on any type that is declared in you package, not just struct types.

```go
type MyFloat float64
func (f MyFloat) Abs() float64 {}
```

However, you cannot define a method on a type from another package.

**methods with pointer receivers**

methods can be associated with a named type or a pointer to a named type.

there are two reasons to use a pointer receiver.

* to avoid copying the value on each method call
* the method can modify the value that it receiver pointer to.

