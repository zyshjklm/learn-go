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

