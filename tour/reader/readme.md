### Reader

the `io` package spcifies the `io.Reader` interface, which represents the read end of a stream of data.

the `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte)  (n int, err error)
```

read populate(人口，-> 填入) the given byte slice b with data and returns the number of bytes populated and an error value.

it returns an `io.EOF` error when the stream ends.

