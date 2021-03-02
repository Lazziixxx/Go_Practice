//一个顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 创建一个服务端  监听tcp端口上的连接
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept() //会持续进行监听 只有当新的连接被创建 才会返回一个net.Conn对象
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

//处理客户端连接的函数
func handleConn(c net.Conn) {
	//向客户端的写操作失败时 服务侧主动关闭net.Conn 回到main函数 等待下一个连接请求
	defer c.Close()
	for {
		//由于net.Conn实现了io.Writer接口
		//服务端不停的获取当前时刻 并写入到客户端
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			//客户端如果主动断开连接 写入失败
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

/*TEST 1*/
/*
使用nc工具 创建网络连接 使当前客户端得到响应
nc64.exe locallhost 8000
 */

/*TEST 2*/
/*
使用go中的net.Dial创建一个Tcp连接
*/

/*TEST 3*/
/*
将handleConn(conn) 修改为go handleConn(conn)  可以同时像多个客户端写入数据
*/
