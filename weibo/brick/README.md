# 微博平台性能挑战赛

## 规则

## V1版本：简单UDP应答，单进程，单线程，Channel通信

### 性能情况
```
Connecting to server at  127.0.0.1:1200
Connected to server at  127.0.0.1:1200
Begin to transfer brick
Transfer OK
receive =  785825619  Bytes
elapsed =  8m54.883491135s  us
bandwidth =  1.4010933641345003  MB/s
```
### 读取文件性能
```
receive =  1063164745  Bytes
elapsed =  9.776157985s  us
bandwidth =  103.71281979800168  MB/s
```
```
SSD benchmark, showing about 230 MB/s reading speed (blue), 210 MB/s writing speed (red) and about 0.1 ms seek time (green), all independent from the accessed disk location.

from http://en.wikipedia.org/wiki/Solid-state_drive
```

### Zero Copy File Transfer

通过GO的io.Copy方法进行大文件快速传输，可以做到内存零拷贝
io.Copy方法优先通过src的WriteTo方法进行拷贝，若不存在则通过dst的ReaderFrom方法进行拷贝，若都不存在则通过内存进行拷贝
GO语言中的bufio实现了WriteTo方法，底层是通系统调用sendFile实现的

```
$ go test -benchmem -bench IoCopy

BenchmarkIoCopy client start reciving.....
receive =  1063164745  Bytes
receive elapsed =  3.065908458s
receive bandwidth =  330.7055397461251  MB/s
       1    3066049413 ns/op    1139665176 B/op 10275724 allocs/op
       ok   github.com/feiyang21687/golang/weibo/brick/benchmark    3.074s

```
### 大文件读取

```
$ go test -benchmem -bench ReadLargeFile

BenchmarkReadLargeFile  receive =  1063164745  Bytes
elapsed =  2.174487509s
bandwidth =  466.27672369632364  MB/s
       1    2174919710 ns/op        4840 B/op         18 allocs/op
       ok   github.com/feiyang21687/golang/weibo/brick/benchmark    2.191s
```

