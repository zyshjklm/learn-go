## lect12 interface



#### interface

* 接口是一个或多个方法签名的集合

* 只要某个类型拥有该接口的所有方法签名，即实现了该接口

* 接口只有方法声明，没有实现

* 将对象赋值给接口时，会发生拷贝，这样将不能修改复制品的状态

* 当接口的对象和类型都是nil时，接口才是nil

* 空接口可以做为任何类型数据的容器。




#### 类型断言

* 通过类型断言的ok, pattern，可以断送接口中的数据类型


* 使用switch, type则可对空接口进行比较全面的类型判断

  ​
#### 接口转换
* 可以将拥有超集的接口转换为子集的接口


在