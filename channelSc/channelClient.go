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
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func()  {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin) //标准输入最后需要加ctrl+z 标志结束输入 否则主goroutine会一直等待
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer,  src io.Reader) {
		if _, err := io.Copy(dst, src) ; err != nil{
			log.Fatal(err)
		}
}

/*
TEST1:
当输入end时，服务侧关闭连接 io.Copy(os.Stdout, conn) 客户端结束读取 打印done 但是主goroutine无法退出 因为stdin没有收到结束符
当输入ctrl+z时 主goroutine关闭对conn的写操作 阻塞等待Channel done获取信号 
同时 服务侧对conn的读操作被关闭 不再scan内容并关闭服务侧的写操作，此时，客户端感知对conn的读操作关闭，打印done并信号到channel，主goroutine读到信号，结束程序
*/
