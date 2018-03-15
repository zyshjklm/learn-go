

## homework



示范写协程池

* 先写一个printURL函数。给定url，抓取url并打印状态
* 定义一个work函数，接收ch中来的url。因此参数是work(ch chan string)
  * 主协程创建chan，并写入数据
  * work获取chan并处理
* ​

