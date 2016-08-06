
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



## ch2/popcountTest

```shell
cd learn-go/gopl/ch2/popcountTest

go run main.go
1: 1
3: 2
7: 3
15: 4
31: 5
63: 6
127: 7
255: 8
511: 9
1023: 10
2047: 11
4095: 12
8091: 10
65535: 16


go test -cpu=4 -bench=. popcount_test.go
testing: warning: no tests to run
PASS
BenchmarkPopCount-4          	200000000	         6.65 ns/op
BenchmarkBitCount-4          	500000000	         3.39 ns/op
BenchmarkPopCountByClearing-4	50000000	        29.0 ns/op
BenchmarkPopCountByShifting-4	20000000	       122 ns/op
ok  	command-line-arguments	8.113s

```

