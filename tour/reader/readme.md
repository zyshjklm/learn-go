### Reader

the `io` package spcifies the `io.Reader` interface, which represents the read end of a stream of data.

the `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte)  (n int, err error)
```

read populate(人口，-> 填入) the given byte slice b with data and returns the number of bytes populated and an error value.

it returns an `io.EOF` error when the stream ends.



```shell
cd ~/_go/src/golang.org/x/
git clone https://github.com/golang/tour
ls -lh
#total 0
#drwxr-xr-x  31 user  staff   1.0K  9 10 17:49 net
#drwxr-xr-x  23 user  staff   782B 10  3 17:00 tour

```

how to generate a err info:

```go
err := fmt.Errorf("Buffer is not long enough")
```



**Rot13Reader**

a common pattern is an io.Reader that wraps another `io.Reader`, modifying the stream in some way.

Rot13:

refer :  https://en.wikipedia.org/wiki/ROT13



tips:

the default length and capacity of a slice in go1.6 is 32768 byte!!!

so :

```go
fmt.Println("cap:", cap(b))
// the loop will be 32768 times.
for i, c := range b {
  	b[i] = rot.cryptRot13(b[i])
}

// recommend: n is the number read returned by rot.r.Read()
for i := 0; i < n; i++ {
    b[i] = rot.cryptRot13(b[i])
}
```

