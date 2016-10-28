## Function

### 1 Multiple return values

Functions and methods can return multiple values. This form can be used to improve on a couple of clumsy(笨拙的) idioms in C program: 

* in-band error returns such as -1 for EOF
* modifying an argument passed by address

**1 error for EOF**

In C, a write error is signaled by a negative count with the error code secreted away in a volatile location. In Go, write can return a count and a error:

"Yes, you wrote some bytes but not all of them because filled the device".

```go
func (file *File) Write(b []byte) ( n int, err error)
// it return the number of bytes written, 
// and a non-nil error when n != len(b)
```

**2 pointer**

A similar approach obviates(消除，避免) the need to pass a pointer to a return value to simulate a reference parameter. 

A function to grab a number from a position in a byte slice, return the number and the next position.

see example: grabNum.go



### 2 Named result parameters

the return of a Go function can be given names and used as regular variables. When named, they are initialized to the zero value for their types when the function begins.

The name are not mandatory(强制的) but they can make code shorter and cleaner.

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
  	for len(buf) > 0; err == nil {
      	var numRead int
      	numRead, err = r.Read(buf)
      	n += numRead
	    buf = buf[numRead:]
  	}
  	return	// means return n, err
}
```

### 3 Defer

defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns.

The canonical example are unlocking a mutex or closing a file.

```go
f, err := os.Open(filename)
if err != nil {
  	return "", err
}
defer f.Close()
```

The arguments to the deferred function are evaluated when the defer executes, not when the `call` executes. 

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
// Output:
// 4 3 2 1 0
```

Deferred functions are executed in LIFO order.

example of use defer to trace function execution: **deferTrace.go**

defer is function-based.


