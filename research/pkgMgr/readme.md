## golang包管理工具

2018.04.30

### 综述

了解一下相关工具。网上资料综述。

* https://gocn.io/question/9
  * 主要推荐的有glide, dep, godep, govender
  * 其中astaxie列出了常见的vendor工具
* vgo
  * 2018年2月Russ Cox设计了新的管理工具。带版本的go



### 版本管理方式

* 《Go的新玩具 vgo》https://zhuanlan.zhihu.com/p/33926171
  * 语义版本 （Semantic Versioning）
    * 典型代表是dep。通过tag来表明依赖的具体版本
  * 导入版本 ( Import Versioning ) 
    *  典型代表 gopkg。大版本，方便修复bug。

### 文件管理方式

* 将所依赖的包文件全放在vendor目录
  * 代码库变大。克隆变慢，好处是下载了即可编译。独立性好。
* 将所依赖的包版本记录到json/toml/yaml文件中
  * 描述文件记录包的commitID。没保留原库的内容。可能原库丢失。 



上述两种分类，只是比较宽泛的说明。还需要处理更细致更复杂场景下的依赖管理。



### 具体信息



* https://github.com/Masterminds/glide
  * 6576 star 比较推荐
* https://github.com/golang/dep
  * 8796 start 比较推荐。亲生的。官方出的实验性工具，可用于生产级
* https://github.com/tools/godep
  * 5335 star。次推荐。
* https://github.com/kardianos/govendor
  * 3335 star。次推荐。
* https://github.com/gogradle/gogradle
  * 421 star，了解即可。
  * 借助了一个自动化构建工具Gradle。因此依赖java1.8



如上，glide应该是现有最好用的，dep是亲生的。两都应该是目前最佳工具。但需要二选一。





