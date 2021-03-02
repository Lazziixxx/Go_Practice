package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err) // log fatal包 defer函数不会执行 打印日志后结束程序退出
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)
	// os stdout本質上是os.File*类型 基于os.File*类型有实现Read和Write接口
}

func mustCopy(dst io.Writer, src io.Reader) {
		if _, err := io.Copy(dst , src); err != nil {
			log.Fatal(err)
		}
}

//使用net.Dial创建了一个Tcp连接作为客户端接收clock服务端的写入数据
//如果创建多个客户端 同一时间只有一个客户端能够输出服务端的传输数据
