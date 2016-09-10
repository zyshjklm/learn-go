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

```
