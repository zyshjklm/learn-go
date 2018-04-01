

### sqlite3



```shell
 sqlite3 database/user.db
SQLite version 3.14.0 2016-07-26 15:17:14
Enter ".help" for usage hints.

### old ops 
sqlite> delete from user where name = 'admin';
sqlite> insert into user (name, age, note, isadmin, password) values ('admin', 20, 'i am admin', 1, );
sqlite> select * from user;                                                                                            2|admin|21232f297a57a5a743894a0e4a801fc3|20|i am admin|1


### new 
sqlite> .tables
sqlite> CREATE TABLE user(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, password TEXT NOT NULL, note TEXT NOT NULL, isadmin SMALLINT NOT NULL);
.tables                                                                                                        user

# insert
sqlite> insert into user (name, password, note, isadmin) values ('admin', '21232f297a57a5a743894a0e4a801fc3', 'i am admin', 1);

sqlite> select * from user;
1|admin|21232f297a57a5a743894a0e4a801fc3|i am admin|1

```



### 功能

实现了增删查改。

实现了静态文件服务器。

有设置session。

没有设置登陆成功、退出等按钮。

