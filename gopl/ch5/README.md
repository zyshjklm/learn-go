## gopl chapter 5

#### How to add a package from Internet

for example, chapter 5 need "golang.org/x/net/html".

because of the GFW, you need git clone the src code.

```shell

$GOPATH

mkdir -p $GOPATH/src/golang.org/x/
cd $GOPATH/src/golang.org/x/

git clone https://github.com/golang/net

# now, you can use "golang.org/x/net/html"

## update, just use 
go get https://github.com/golang/net

```

#### findlink

```shell

go build ../ch1/fetch/fetchall.go
# fetchall

go build findlinks1/main.go

./fetchall https://golang.org > go.html

cat go.html | ./main

```
