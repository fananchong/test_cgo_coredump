# test_cgo_coredump
cgo 中 crash ，捕获 go 堆栈、 c 堆栈

两个堆栈无法同时捕获，但是可以分别捕获 go 调用栈、 c 调用栈


## 捕获 C++ 错误，并做善后处理

例子参见：
- [main2.go](main2.go)
- [catch_except.go](catch_except.go)

编译使用 
```shell
make all2
```

测试使用
```shell
./main2
```


## go 调用栈

cgo 异常必 crash ， go 进程最后会输出所有堆栈信息。

因此捕获这些信息即可。

通常这时已经无法使用程序内 log 模块，因此，stderr 重定向到文件即可，类似：

```bash
./main > 1.log 2>&1
```

或者

```bash
./main &2> 1.log
```


本例子输出结果会是这样子：

```bash
main: test.c:9: fn2: Assertion `1 == 2' failed.
test_crash from C and here the str is from Go: From GolangArgument deadbeefSIGABRT: abort
PC=0x7f0f0b860428 m=0 sigcode=18446744073709551610

goroutine 0 [idle]:
runtime: unknown pc 0x7f0f0b860428
stack: frame={sp:0x7ffc4d97cb58, fp:0x0} stack=[0x7ffc4d17dec8,0x7ffc4d97cf00)
00007ffc4d97ca58:  00007ffc4d97ccb8  0000000000c1a560 
00007ffc4d97ca68:  00007f0f0b8a17fc  00000000fbad8000 
00007ffc4d97ca78:  0000000000c1a560  0000000000c1a560 
00007ffc4d97ca88:  00007f0f0b9222e9  0000000000000001 
00007ffc4d97ca98:  00007f0f0b8a3bff  00007f0f0bbf0620 
00007ffc4d97caa8:  000000000000004b  0000000000c19550 
00007ffc4d97cab8:  0000000000000000  0000000000000000 
00007ffc4d97cac8:  00007f0f0b8a5409  0000000000000000 
00007ffc4d97cad8:  00007f0f0bbf0620  0000000000000000 
00007ffc4d97cae8:  00007f0f0c02a700  0000000000000000 
00007ffc4d97caf8:  00007f0f0b8a7196  0000000000000000 
00007ffc4d97cb08:  0000000000000000  ffffffffffffffff 
00007ffc4d97cb18:  0000000000000000  0000000000000000 
00007ffc4d97cb28:  00007f0f0c036000  00000000004c4dab 
00007ffc4d97cb38:  0000000000000009  00000000004c4def 
00007ffc4d97cb48:  000000000000001f  0000000000000200 
00007ffc4d97cb58: <00007f0f0b86202a  0000000000000020 
00007ffc4d97cb68:  0000000000000000  0000000000000000 
00007ffc4d97cb78:  0000000000000000  0000000000000000 
00007ffc4d97cb88:  0000000000000000  0000000000000000 
00007ffc4d97cb98:  0000000000000000  0000000000000000 
00007ffc4d97cba8:  0000000000000000  0000000000000000 
00007ffc4d97cbb8:  0000000000000000  0000000000000000 
00007ffc4d97cbc8:  0000000000000000  0000000000000000 
00007ffc4d97cbd8:  0000000000000000  0000000000000004 
00007ffc4d97cbe8:  0000000000000000  0000000000000000 
00007ffc4d97cbf8:  00007ffc4d97cbc0  00000016839455a4 
00007ffc4d97cc08:  00007f0f0c036000  00000000004c4dab 
00007ffc4d97cc18:  0000000000000009  00000000004c4def 
00007ffc4d97cc28:  000000000000001f  0000000000000200 
00007ffc4d97cc38:  00007f0f0b8af53c  00007f0f0b9bb250 
00007ffc4d97cc48:  00007f0f0b9be7a0  0000000000000000 
runtime: unknown pc 0x7f0f0b860428
stack: frame={sp:0x7ffc4d97cb58, fp:0x0} stack=[0x7ffc4d17dec8,0x7ffc4d97cf00)
00007ffc4d97ca58:  00007ffc4d97ccb8  0000000000c1a560 
00007ffc4d97ca68:  00007f0f0b8a17fc  00000000fbad8000 
00007ffc4d97ca78:  0000000000c1a560  0000000000c1a560 
00007ffc4d97ca88:  00007f0f0b9222e9  0000000000000001 
00007ffc4d97ca98:  00007f0f0b8a3bff  00007f0f0bbf0620 
00007ffc4d97caa8:  000000000000004b  0000000000c19550 
00007ffc4d97cab8:  0000000000000000  0000000000000000 
00007ffc4d97cac8:  00007f0f0b8a5409  0000000000000000 
00007ffc4d97cad8:  00007f0f0bbf0620  0000000000000000 
00007ffc4d97cae8:  00007f0f0c02a700  0000000000000000 
00007ffc4d97caf8:  00007f0f0b8a7196  0000000000000000 
00007ffc4d97cb08:  0000000000000000  ffffffffffffffff 
00007ffc4d97cb18:  0000000000000000  0000000000000000 
00007ffc4d97cb28:  00007f0f0c036000  00000000004c4dab 
00007ffc4d97cb38:  0000000000000009  00000000004c4def 
00007ffc4d97cb48:  000000000000001f  0000000000000200 
00007ffc4d97cb58: <00007f0f0b86202a  0000000000000020 
00007ffc4d97cb68:  0000000000000000  0000000000000000 
00007ffc4d97cb78:  0000000000000000  0000000000000000 
00007ffc4d97cb88:  0000000000000000  0000000000000000 
00007ffc4d97cb98:  0000000000000000  0000000000000000 
00007ffc4d97cba8:  0000000000000000  0000000000000000 
00007ffc4d97cbb8:  0000000000000000  0000000000000000 
00007ffc4d97cbc8:  0000000000000000  0000000000000000 
00007ffc4d97cbd8:  0000000000000000  0000000000000004 
00007ffc4d97cbe8:  0000000000000000  0000000000000000 
00007ffc4d97cbf8:  00007ffc4d97cbc0  00000016839455a4 
00007ffc4d97cc08:  00007f0f0c036000  00000000004c4dab 
00007ffc4d97cc18:  0000000000000009  00000000004c4def 
00007ffc4d97cc28:  000000000000001f  0000000000000200 
00007ffc4d97cc38:  00007f0f0b8af53c  00007f0f0b9bb250 
00007ffc4d97cc48:  00007f0f0b9be7a0  0000000000000000 

goroutine 1 [syscall]:
runtime.cgocall(0x47e650, 0xc000043f60, 0xb)
        /usr/local/go/src/runtime/cgocall.go:133 +0x5b fp=0xc000043f30 sp=0xc000043ef8 pc=0x40431b
main._Cfunc_test_crash(0xc19050)
        _cgo_gotypes.go:59 +0x41 fp=0xc000043f60 sp=0xc000043f30 pc=0x47e4f1
main.main()
        /home/fananchong/test_cgo_coredump/main.go:20 +0x44 fp=0xc000043f88 sp=0xc000043f60 pc=0x47e5f4
runtime.main()
        /usr/local/go/src/runtime/proc.go:203 +0x212 fp=0xc000043fe0 sp=0xc000043f88 pc=0x431bc2
runtime.goexit()
        /usr/local/go/src/runtime/asm_amd64.s:1373 +0x1 fp=0xc000043fe8 sp=0xc000043fe0 pc=0x45c3d1

rax    0x0
rbx    0x7f0f0c036000
rcx    0x7f0f0b860428
rdx    0x6
rdi    0x3e92
rsi    0x3e92
rbp    0x4c4dab
rsp    0x7ffc4d97cb58
r8     0xc19010
r9     0xfefefefefefefe00
r10    0x8
r11    0x202
r12    0x9
r13    0x4c4def
r14    0x1f
r15    0x200
rip    0x7f0f0b860428
rflags 0x202
cs     0x33
fs     0x0
gs     0x0
```

如果进程协程数很多，则可以搜索关键字`cgocall`定位日志


##### dlv

注意 dlv 只能排查 go 这边的代码，因此对于 cgo 调用栈目前是无能为力的

dlv 相关用法，可以参考： [https://blog.csdn.net/KentZhang_/article/details/84925878](https://blog.csdn.net/KentZhang_/article/details/84925878)

安装： `go get github.com/derekparker/delve/cmd/dlv` 会安装到 GOPATH/bin 目录下


## c 调用栈

1. `ulimit -c unlimited` 确保可以生产 core

2. `env GOTRACEBACK=crash ./main` 按这种方式执行程序

3. 会生产 core 文件，可以使用 gdb 调试


类似以下函数栈：

```bash
(gdb) t 2
[Switching to thread 2 (Thread 0x7f9f14feb700 (LWP 16123))]
#0  runtime.usleep () at /usr/local/go/src/runtime/sys_linux_amd64.s:146
146             RET
(gdb) bt
#0  runtime.usleep () at /usr/local/go/src/runtime/sys_linux_amd64.s:146
#1  0x0000000000442b7e in runtime.sighandler (sig=6, info=0xc0000094b0, ctxt=0xc000009380, gp=0x73f1a0 <runtime.g0>)
    at /usr/local/go/src/runtime/signal_unix.go:642
#2  0x0000000000442479 in runtime.sigtrampgo (sig=6, info=0xc0000094b0, ctx=0xc000009380)
    at /usr/local/go/src/runtime/signal_unix.go:444
#3  0x000000000045e2b3 in runtime.sigtramp () at /usr/local/go/src/runtime/sys_linux_amd64.s:389
#4  <signal handler called>
#5  0x00007f9f14821428 in __GI_raise (sig=sig@entry=6) at ../sysdeps/unix/sysv/linux/raise.c:54
#6  0x00007f9f1482302a in __GI_abort () at abort.c:89
#7  0x00007f9f14819bd7 in __assert_fail_base (fmt=<optimized out>, assertion=assertion@entry=0x4c4dab "1 == 2", 
    file=file@entry=0x4c4da4 "test.c", line=line@entry=9, function=function@entry=0x4c4def <__PRETTY_FUNCTION__.2556> "fn2")
    at assert.c:92
#8  0x00007f9f14819c82 in __GI___assert_fail (assertion=0x4c4dab "1 == 2", file=0x4c4da4 "test.c", line=9, 
    function=0x4c4def <__PRETTY_FUNCTION__.2556> "fn2") at assert.c:101
#9  0x000000000047edd0 in fn2 (arg=0x7ffe75c08c20 "deadbeef") at test.c:9
#10 0x000000000047ee0f in fn1 (arg=1092) at test.c:17
#11 0x000000000047ee3e in test_crash (str=0x266a050 "From Golang") at test.c:23
---Type <return> to continue, or q <return> to quit---
#12 0x000000000045bb20 in runtime.asmcgocall () at /usr/local/go/src/runtime/asm_amd64.s:655
#13 0x000000c000054120 in ?? ()
#14 0x00000000004b1010 in func.* ()
#15 0x0000000000000010 in ?? ()
#16 0x0000000000491a80 in type.* ()
#17 0x0000000000491a80 in type.* ()
#18 0x0000000000000110 in ?? ()
#19 0x000000c000000180 in ?? ()
#20 0x000000000045a346 in runtime.systemstack () at /usr/local/go/src/runtime/asm_amd64.s:370
#21 0x0000000000434340 in ?? () at <autogenerated>:1
#22 0x000000000045a1d4 in runtime.rt0_go () at /usr/local/go/src/runtime/asm_amd64.s:220
#23 0x000000000047ee50 in test_crash (str=0x7ffe75409d98 <error: Cannot access memory at address 0x7ffe75409d98>) at test.c:25
#24 0x000000000045a1db in runtime.rt0_go () at /usr/local/go/src/runtime/asm_amd64.s:225
#25 0x0000000000000001 in ?? ()
#26 0x00007ffe75c08ed8 in ?? ()
#27 0x0000000000000001 in ?? ()
#28 0x00007ffe75c08ed8 in ?? ()
---Type <return> to continue, or q <return> to quit---
#29 0x0000000000000000 in ?? ()
(gdb) 
```


## 其他 - 脚本获取 cgo 所在线程


```vim
#!/bin/bash

(gdb ./main ./core > a.log )<< GDBEOF
thread apply all bt
GDBEOF

n=`grep "runtime.rt0_go" a.log  | awk -F' ' '{printf $1 }'`
n=${n:1}
line=`grep -n "runtime.rt0_go" a.log | awk -F':' '{printf $1 }'`

no=$((line-n-2))
sed -n $no','$no'p' a.log
```

输出类似：

```vim
fananchong@fananchong-ubuntu:~/test_cgo_coredump$ ./g.sh 

warning: Unexpected size of section `.reg-xstate/2278388' in core file.

warning: Unexpected size of section `.reg-xstate/2278388' in core file.
warning: File "/usr/local/go/src/runtime/runtime-gdb.py" auto-loading has been declined by your `auto-load safe-path' set to "$debugdir:$datadir/auto-load".
Thread 4 (Thread 0x7f60186f1740 (LWP 2278383)):
```

最后一行`Thread 4 ...`就是 cgo 所在线程号


## 其他 - c 代码中访问空指针，导致堆栈被破坏

在 cgo 内捕获异常，可以避免该问题。详细参考 [catch_except.go](catch_except.go)


## 其他 - breakpad

在 cgo 中使用 breakpad 暂时有 2 个问题：
- 调用栈，显示不了符号（可能我哪里打开方式不对？）
- 回调 golang 代码会提示：
  ```shell
  Dump path: ./2bfc565f-cf9e-49b2-89ef789c-f3792904.dmp
  fatal: morestack on g0
  Trace/breakpoint trap(吐核)
  ```
  
  golang 回调代码无法正常被调用




