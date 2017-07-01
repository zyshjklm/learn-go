



Return code

```shell

go run return.go
open abcd.txt: no such file or directory
exit status 2

go build return.go

./return
open abcd.txt: no such file or directory

echo $?
2
```

