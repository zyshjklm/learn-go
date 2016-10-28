## Function

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

















