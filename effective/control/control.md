## Control structures



The control structure of Go are related to those of C but differ in important ways.

There is no do or while loop.

### 1 If

```go
if x > 0 {
  	return y
}
```

Mandatory(强制的) braces encourage writing simple if statement on multiple lines.

if accept an initialization statement.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

This is an example of a common situation where code must guard against a sequence of error conditions. 

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### 2 redeclaration and reassignment

the previous example demonstrates a detail of how the `:=` short declaration form works.

```go
f, err := os.Open(name)
d, err := f.Stat()
```

the  `err`  is declared by the first statement, but only re-assigned in the second.

in a `:=` declaration, a variable v may appear even if it has already been declared, **provide**:

* the declaration is in the same scope as the existing declaration of v.
* the value in the initialization is assignable to v, and
* there is at least one other variable in the declaration that is being declared anew(全新的).

This property is pure pragmatism 实用主义.

Also, the scope of function parameters and return values in the same as the function body.

### 3 for

usage:

```go
for init; condition; post {}	// Like a C for
for condition {}	// Like a C while
for {}

sum := 0
for i := 0; i < 10; i++ {
  	sum += i
}
```

looping over an array, slice, string or map, or reading from channel. `range` clause.

```go
for key, value := range oldMap {
  	newMap[key] = value
}

// drop second
for key := range m {
  	if key.expired() {
    	delete(m, key)	
  	}
}

// drop first
sum := 0
for _, value := range array {
  	sum += value
}
```

If you only need the first item in the range (the key or index), drop the second as above.

for strings, the range breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune **U+FFFD**. `rune` is Go terminology专门术语 for a single Unocode code point.

example: for_string.go

Finally, Go has no comma operator and ++ and — are statement not expressions.

### 4 switch

The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found.
```go
func unhex(c byte) byte {
  	switch {
    case '0' <= c && c <= '9':
    	return c - '0'
    case 'a' <= c && c <= 'f':
      	return c - 'a' + 10
    case 'A' <= c && c <= 'F':
      	return c - 'A' + 10
  	}
  	return 0
}
```

There is no automatic fall through, but cases can be presented in comma-separated lists.

```go
func shouldEscape(c byte) bool {
  	switch c {
    case ' ', '?', '&', '=', '+', '%':
      	return true
  	}
  	return false
}
```



#### 4.1 switch and break

`break` statements can be used to terminate a switch early.

```go
Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break	// terminate switch
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				break Loop	// terminate Loop 
			}
			if validateOnly {
				break	// terminate switch
			}
			size = 2
			update(src[n] + src[n+1]<<shift)
		}
	}
```

a example of switch: compare.go

#### 4.2 type switch

A switch can also be used to discover the dynamic type of an interface variable.

```go
var t interface{}
t = funcOfSomeType()

switch t := t.(type) {	
  // declaring a new var t with the same name of outside 
  // but a different type in each case.
default: 
  	fmt.Printf("unexpected type %T\n", t)	// %T print whatever type t has
case bool:
  	fmt.Printf("boolean %t\n", t)
case int:
  	fmt.Printf("int %t\n", t)
case *bool:
  	fmt.Printf("pointer of boolean %t\n", *t)
case *int:
  	fmt.Printf("pointer of int %t\n", *t)
}
```

details in example: switchType.go

