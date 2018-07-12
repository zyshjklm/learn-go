## golang包管理

### 一、综述

了解一下相关工具。网上资料综述。

- https://gocn.io/question/9
  - 主要推荐的有glide, dep, godep, govender
  - 其中astaxie列出了常见的vendor工具
- vgo
  - 2018年2月Russ Cox设计了新的管理工具。带版本的go



### 二、版本管理方式

- 《Go的新玩具 vgo》https://zhuanlan.zhihu.com/p/33926171
  - 语义版本 （Semantic Versioning）
    - 典型代表是dep。通过tag来表明依赖的具体版本
  - 导入版本 ( Import Versioning ) 
    - 典型代表 gopkg。大版本，方便修复bug。



### 三、文件管理方式

- 将所依赖的包文件全放在vendor目录
  - 代码库变大。克隆变慢，好处是下载了即可编译。独立性好。
- 将所依赖的包版本记录到json/toml/yaml文件中
  - 描述文件记录包的commitID。没保留原库的内容。可能原库丢失。 

上述两种分类，只是比较宽泛的说明。还需要处理更细致更复杂场景下的依赖管理。



### 四、实际管理工具信息

主要是github上的star比较。

- https://github.com/Masterminds/glide
  - 6576 star 比较推荐
- https://github.com/golang/dep
  - 8796 start 比较推荐。亲生的。官方出的实验性工具，可用于生产级
- https://github.com/tools/godep
  - 5335 star。次推荐。
- https://github.com/kardianos/govendor
  - 3335 star。次推荐。
- https://github.com/gogradle/gogradle
  - 421 star，了解即可。
  - 借助了一个自动化构建工具Gradle。因此依赖java1.8

如上，glide应该是现有最好用的，dep是亲生的。两都应该是目前最佳工具。



### 五、vendor目录

#### 5.1 不使用包管理

```shell
# go get github.com/jkak/test
# ls -l $GOPATH/src/github.com/jkak/test/mytest/mytest.go

# mkdir nopkg && cd nopkg
# vim main1.go
# go run main1.go
hello golang!
```



#### 5.2 使用vendor目录

测试使用vendor目录。将`go get`下载的包文件，移到本地的vendor目录下。依然可以运行。也就是，go在编译时，如果发现有vendor目录，则会优先在vendor目录下寻找依赖包。

```shell
# mkdir -p vendor/github.com/jkak/
# mv $GOPATH/src/github.com/jkak/test vendor/github.com/jkak/
# ls -lh vendor/github.com/jkak/test/mytest
# ls -l $GOPATH/src/github.com/jkak/test
ls: ~/src/github.com/jkak/test: No such file or directory

# go run main1.go
hello golang!
```

需要注意的是，vendor下的目录结构，需要与`$GOPATH/src/`目录一致。





### 六、glide

环境安装

```shell
go get -u github.com/Masterminds/glide

cd $GOPATH/src/github.com/Masterminds/glide
go install

glide
```

测试使用示例：

- github.com/jkak/test/mytest



#### 6.1 install

```shell
# cd ../
# mkdir test1-glide
# cp nopkg/main1.go test1-glide/
# cd test1-glide/

# glide init
glide init
[INFO]	Generating a YAML configuration file and guessing the dependencies
[INFO]	Attempting to import from other package managers (use --skip-import to skip)
[INFO]	Scanning code to look for dependencies
[INFO]	--> Found reference to github.com/jkak/test/mytest
[INFO]	Writing configuration file (glide.yaml)
[INFO]	Would you like Glide to help you find ways to improve your glide.yaml configuration?
[INFO]	If you want to revisit this step you can use the config-wizard command at any time.
[INFO]	Yes (Y) or No (N)?
Y
[INFO]	Looking for dependencies to make suggestions on
[INFO]	--> Scanning for dependencies not using version ranges
[INFO]	--> Scanning for dependencies using commit ids
[INFO]	Gathering information on each dependency
[INFO]	--> This may take a moment. Especially on a codebase with many dependencies
[INFO]	--> Gathering release information for dependencies
[INFO]	--> Looking for dependency imports where versions are commit ids
[INFO]	No proposed changes found. Have a nice day.

# ls
glide.yaml main1.go
# cat glide.yaml
package: github.com/jkak/learn-go/reboot3/pkgMgr/test1-glide
import:
- package: github.com/jkak/test
  subpackages:
  - mytest

# glide install
[INFO]	Lock file (glide.lock) does not exist. Performing update.
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching github.com/chenchao1990/myClass
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/chenchao1990/myClass
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 1 dependencies.

#  go run main1.go
hello golang!

# ls -l vendor/github.com/jkak/test/mytest/mytest.go
```



#### 6.2 update

先在本地仓库修改依赖包。

```shell
# cd $GOPATH/src/github.com/jkak/test/mytest
# vim mytest.go #### to add Test()

# git commit -m 'add Test() in mytest/mytest.go'
[master 0c89b4a] add Test() in mytest/mytest.go
```

上述代码先不push到github。直接到test1-glide目录进行update操作。

```shell
# cd $GOPATH/src/github.com/jkak/learn-go/reboot3/pkgMgr/test1-glide
glide up
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for github.com/jkak/test
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/jkak/test
[INFO]	Replacing existing vendor dependencies
[INFO]	Versions did not change. Skipping glide.lock update.
[INFO]	Project relies on 1 dependencies.
```

日志显示，并没有执行任何修改。



下面提交test代码

```shell
# cd $GOPATH/src/github.com/jkak/test/mytest
# git push
Counting objects: 4, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 431 bytes | 0 bytes/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To https://github.com/jkak/test
   a5c916f..0c89b4a  master -> master
```

回到test1-glide

```shell
# cd $GOPATH/src/github.com/jkak/learn-go/reboot3/pkgMgr/test1-glide

# glide up
[INFO]	Downloading dependencies. Please wait...
[INFO]	--> Fetching updates for github.com/jkak/test
[INFO]	Resolving imports
[INFO]	Downloading dependencies. Please wait...
[INFO]	Setting references for remaining imports
[INFO]	Exporting resolved dependencies...
[INFO]	--> Exporting github.com/jkak/test
[INFO]	Replacing existing vendor dependencies
[INFO]	Project relies on 1 dependencies.

# vim main1.go	#### add mytest.Test()

# go run main1.go
hello golang!
golang test func
```



其它参考：

* https://segmentfault.com/a/1190000005929355
* https://www.jianshu.com/p/7751affc933e



### 七、godep

安装工具

```shell
# go get github.com/tools/godep
# which godep
```

测试环境

```shell
# cd ../
# mkdir test2-godep
# cp -v test1-glide/main1.go test2-godep
# cd test2-godep 
```

使用godep

```shell
# rm -rf $GOPATH/src/github.com/jkak/test/
# ls -lh $GOPATH/src/github.com/jkak/test/	### 不存在了

# godep get
# ls -lh $GOPATH/src/github.com/jkak/test/
total 8
-rw-r--r--  1 song  staff   33 Jul 12 23:48 README.md
drwxr-xr-x  3 song  staff  102 Jul 12 23:48 mytest
drwxr-xr-x  3 song  staff  102 Jul 12 23:48 server

# ls
main1.go
# godep save

# ls -lh
total 8
drwxr-xr-x  4 song  staff   136B Jul 12 23:50 Godeps
-rw-r--r--  1 song  staff   104B Jul 12 23:46 main1.go
drwxr-xr-x  3 song  staff   102B Jul 12 23:50 vendor
# ls -lh vendor/github.com/jkak/test
total 0
drwxr-xr-x  3 song  staff   102B Jul 12 23:50 mytest

# go run main1.go
hello golang!
golang test func
```

如上，只引用了test中的mytest包。

总结godep相关操作：

* godep get 是将包存在到$GOPATH/src/目录
* godep save将$GOPATH/src/下的包保存到本地的vendor。
* godep restore可以还源$GOPATH/src/下的没有的目录。
  * 比如$GOPATH/src/下没有github.com/jkak/test时，
  * 就可以该命令将vendor下的库用来恢复到$GOPATH/src/下。
* godep update用来更新。
* 不要随意去清除包里的`.git`信息，不然会影响godep的使用。
  * 如果直接在$GOPATH/src/下相关的包里进行了修改，但没有保存并commit。
  * 则godep save会提示出现了dirty数据。


