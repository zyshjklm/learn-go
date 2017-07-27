

### 下载 

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
