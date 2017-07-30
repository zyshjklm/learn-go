## miniFtp.go

client从服务器获取文件：`GET XXXX.file.name`

```shell
#### server 
# go run miniFtp.go
2017/07/30 23:27:17 root: ./
2017/07/30 23:27:17 root: ./
2017/07/30 23:27:23 cmd:GET, name:/a.txts
2017/07/30 23:27:23 open .//a.txts: no such file or directory
2017/07/30 23:27:30 cmd:GET, name:/a.txt

#### client
# echo 'GET /a.txts' | nc localhost 8021
open .//a.txts: no such file or directory%                                                                                        

# echo 'GET /a.txt' | nc localhost 8021
hello golang!
```

如果使用前面实现的client

```shell
# go build -o ./client ../server/client_5file.go
# ./client
#GET /a.txt
2017/07/30 23:30:03 write size: 23
2017/07/30 23:30:03 return content:
hello golang!
```



