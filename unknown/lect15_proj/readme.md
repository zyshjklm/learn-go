## lect15 project

创建项目。

被调用包：

```shell
cd $GOPATH

cd src/github.com/learn-go/unknown/
mkdir -p lect15_proj/{goTest, call}

cd lect15_proj/GoTest/
vi goTest.go
# add program

go run
# compile to:
# $GOPATH/pkg/darwin_amd64/github.com/learn-go/unknown/lect15_proj/goTest.a
```



主调函数包：

```shell
cd lect15_proj/call/
vi main.go
# add program

go run main.go
#Hello from goTest.
#Hi from goTest.
```



