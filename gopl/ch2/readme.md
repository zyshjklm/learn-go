
## ch2/echo

```shell
# build
go build github.com/learn-go/gopl/ch2/echo

./echo -n

./echo -s "*" a bc defg
#a*bc*defg

./echo -s " * " a bc defg
a * bc * defg

./echo -s " * " -n a bc defg
a * bc * defg%

# help
./echo -s
flag needs an argument: -s
Usage of ./echo:
  -n	omit trailing newline
  -s string
    	separator (default " ")

./echo -help
Usage of ./echo:
  -n	omit trailing newline
  -s string
    	separator (default " ")

./echo -h
Usage of ./echo:
  -n	omit trailing newline
  -s string
    	separator (default " ")
```

