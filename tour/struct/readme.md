## pointers and structs

refer from : http://tour.golangtc.com/moretypes/1

### 1 pointers

a pointer holds the memory address of a variable.

the type `*T` is a pointer to a `T` value. its zero value is `nil`.

* the & operator generates a pointer to its operand.
* the * operator denotes the pointer's underlying value.

unlick C, Go **has no pointer arithmetic.**

```go
var p *int
var q *byte

i := 43
p = &i
```

example in pointer.go.

#### 2. struct

a struct is a collection of fileds. a `type` declaration does what you'd expect.

struct fields are accessed using a dot.

**pointer to struct:** struct field can be accessed through a struct pointer.

**struct literal**: denotes a newly allocated struct value by listing the values of its fields.

you can list just a subset of fields by using the `Name:` syntax.

the special prefix & returns a pointer to the struct value.

 

