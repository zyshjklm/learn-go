
### test1

直接使用toml，包括常用的key/value，也包括slice，后者需要在配置中使用`[[]]`来表示。
所有字段都需要首字母大写。

如果有使用map，则配置中使用`[]`来表示。 

### test2 

给定义的结果增加tag，比如`toml: "trans_addr"`这样就可以不需要toml中写大写首字母了。

toml文件中的key，直接使用struct中的toml tag的值。

### test3

定义一个sender结构体，将相应的字段都放入结构体中。struct也需要对应一个`[]`。

将step由字符串改为int，表示秒数。



