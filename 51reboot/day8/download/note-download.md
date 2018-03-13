

### 下载 

* 从原始链接进行下载，解析，清洗新链接
* 建立临时目录。使用ioutil.TmpDir()
* 下载
  * 使用http.Get()下载，处理错误，并defer关闭
  * 用filepath.Join(tmpDir, path.Base(url))生成文件名
  * 用os.Create(fullNmae)创建文件名
  * 用io.Copy(fd, resp.Body)将抓取结果写入文件
  * 关闭文件



### 并发下载

将原本的串行下载，改成通过chan来控制的并发下载。

三层函数：

* downloadImgs(urls []string, dir string, conNum int)
  * 入口函数，要下载的包的链接，存储目录，并发数
  * 检查并发数；设置WaitGroup；创建chan，一个缓冲即能运行
  * 根据并发，启动协程：go downWorker(urlChan, dir, &wg)
  * 循环urls，将url放入chan中。然后关闭chan。
  * 等待wg.Wait()
* downWorker(urlChan chan string, dir string, wg *sync.WaitGroup)
  * 通过for range从Chan中读取url，然后调用downloadURL
  * 等for结束后，调用wg.Done()，每个并发的Worker都需要Done 
* downloadURL(url, dir string)
  * resp = http.Get(url)
  * fd = os.Create(fullName)
  * io.Copy(fd, resp.Body)


### 打包

code: 

* 4maketar/4maketar.go
* 4multiDown2Targz/main.go

```go
func Walk(root string, walkFn WalkFunc) error
// Walk walks the file tree rooted at root, calling walkFn for
// each file or directory in the tree, including root.

type WalkFunc func(path string, info os.FileInfo, err error) error
// function called for each file or directory visited by Walk.
// path包括Walk中的root部分，info是path对应的文件信息

type FileInfo interface {}
// A FileInfo describes a file and is returned by Stat and Lstat.

func Rel(basepath, targpath string) (string, error)
// Join(basepath, Rel(basepath, targpath)) 
// is equivalent to targpath itself.
// base = "/a", target = "/a/b/c", => 结果："b/c"
```



基本顺序：

* 使用os.Create(tarName)打开文件，得到fd
* 使用tar.NewWriter(fd)封装Writer接口。得到tr
* 遍历被打包对象，进行打包。

打包使用filepath.Walk函数。Walk的第二个参数是个函数，函数的主要工作：

* 写入tar的FileHeader
* 以读取的方式打开文件
* 判断目录和文件，如果是文件
* 把文件的内容写入到body

匿名函数里，注意name与header.Name的区别，前者有目录，后者没有目录。

原因参考archive/tar#FileInfoHeader的说明。这个函数os.FileInfo方法只返回文件描述符所代表文件的`base name`。header.Name来自info.Name()。

```shell
go run 4maketar/4maketar.go
# 查看文件列表
tar tf img.tar
```



### 压缩

有了打包的功能，添加压缩就比较简单了。对比一下：

```go
// tar of fd
tr := tar.NewWriter(fd)
defer tr.Close()

// tar.gz of fd
compress := gzip.NewWriter(fd)
tr := tar.NewWriter(compress)
defer compress.Close()
defer tr.Close()
```



### 下载不同的标签

code: 5down2Lable/main.go

使用一个map来表示不同的类型。

 

```shell
doc.Find("img").Each(func(i int, s *goquery.Selection) {
下载js
    改为script

下载链接：
    改为a, 获取的地方则是s.Attr("href")


## run 

cd 5down2Lable
go run main.go -w "http://59.110.12.72:7070/golang-spider/img.html"

 0 https://pic1.zhimg.com/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
 1 http://pic4.zhimg.com/v2-40becd4a519329198ecb3807f342fd7b.jpg
 2 http://59.110.12.72:7070/golang-spider/img/a.jpg
 3 http://59.110.12.72:7070/golang-spider/img/b.jpg
 2017/07/27 22:31:40 down: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider275440321/a.jpg
 2017/07/27 22:31:40 down: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider275440321/b.jpg
 2017/07/27 22:31:40 down: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider275440321/v2-40becd4a519329198ecb3807f342fd7b.jpg
 2017/07/27 22:31:41 down: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider275440321/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
 start make tar...: /var/folders/4_/20sh50zs0lg64gphvscqq2gm0000gn/T/spider275440321


 ls -tlr
 total 400
 -rw-r--r--  1 song  staff    3951 Jul 27 22:31 main.go
 -rw-r--r--  1 song  staff  198900 Jul 27 22:31 img.tar.gz

 tar -zxvf img.tar.gz
 x spider275440321
 x spider275440321/a.jpg
 x spider275440321/b.jpg
 x spider275440321/v2-40becd4a519329198ecb3807f342fd7b.jpg
 x spider275440321/v2-58e318de6172810c1b3c7236e8e0dbb4.jpg
```

