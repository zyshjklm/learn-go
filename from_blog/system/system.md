## system info for goroutine

[refer from](http://blog.csdn.net/kjfcpua/article/details/17451507)

### 1 OS and running library

示意图：

###### ![vdso](vdso.png)



VDSO virtual dynamic shared object.随内核发行但不在内核态
简单调用直接由glibc提供，其它由VDSO提供接口并决定后继调用接口
对VDSO，fork走INT80，getuid走syscall.

linux: http://syscalls.kernelgrok.com/
win: http://j00ru.vexillium.org/ntapi/
VDSO: http://lwn.net/Articles/446528/

