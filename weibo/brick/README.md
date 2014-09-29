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



