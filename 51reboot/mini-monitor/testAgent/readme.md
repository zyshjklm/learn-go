

### testAgent

使用

agent1.go

```shell
# go run testAgent/agent1.go
2017/08/19 23:03:51 {"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":5.527638190954774,"timestamp":1503155031}
2017/08/19 23:03:52 {"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":8.5,"timestamp":1503155032}

```



agent2.go

```Shell
# nc -l 127.0.0.1 6000

# go run agent2.go
2017/08/19 23:02:23 write len:218, buf len:218,  status:true
2017/08/19 23:02:23 [{"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":5.486284289276807,"timestamp":1503154943},{"metric":"mem.usage","endpoint":"jungle85","tag":["darwin"],"value":573650124800,"timestamp":1503154943}]
```

需要启动一个监听端口。



agent3.go

```Shell
# nc -l 127.0.0.1 6000

# go run agent3.go
2017/08/19 23:26:10 {"metric":"cpu.usage","endpoint":"jungle85","tag":["darwin"],"value":3.015075376884422,"timestamp":1503156370}
2017/08/19 23:26:10 {"metric":"mem.usage","endpoint":"jungle85","tag":["darwin"],"value":522215424000,"timestamp":1503156370}
```

