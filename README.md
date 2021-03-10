# Go_Practice
Go Simple Demo For Practice

clockSc:
clockClient.go
clockService.go
建立tcp连接，客户端打印服务端的clock信息

echoSc：
echoClient.go
echoService.go
建立tcp连接，客户端输入xxx，服务端按回声的方式返回

channelSc:
channelClient.go
channelService.go
建立tcp连接，通过无缓存的channel + connect读写关闭实现客户端goroutine同步

Performance:
性能分析目录 主要对比分析各个基础操作的最优实现，并探究背后的原理

benchmark：
benchmark 测试方法

pprof：
pprof性能分析方法

