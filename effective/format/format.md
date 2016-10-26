## introduce

to write Go well, it is important to understand its properties and idioms.

This document gives tips for writing clear, idiomatic Go code.

## example

the Go package sources are intended to serve not only as the core library 
but also as examples of how to use the language.

many package contain working, self-contained executable examples 
you can run directly from web site.

## formatting

Formatting issues are the most contentious but the least consequential.
格式化问题是最有争议，却最少达成共识的问题。

gofmt program (also avaliable as go fmt, which operate at the package level
rather than source file level) read a Go program and emits the source in a 
standard style of indentation and vertical alignment, retaining and if 
necessary reformatting comments.

example of formatEg1.go

```shell

cat formatEg1.go

gofmt formatEg1.go > result.go
# see the result.go

# or
go fmt formatEg1.go
# will update the formatEg1.go file as result.go
```

about gofmt, you can try as below:

```shell

# show diffs
gofmt -d formatEg1.go

gofmt -w  formatEg1.go
# instead of :
#   gofmt formatEg1.go > result.go
#   mv result formatEg1.go

gofmt -h

```

formatting details remain:

* use tab for indentation 
* has no line length limit. wrap long line and indent with an extra tab.
* needs fewer parentheses.
  * if, for, switch do not have parentheses
  * operator precedence hierarchy is shorter and cleaner, 
    * like: x<<8 + y<<16, means what the spacing implies

