package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listen , err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

//将客户端传给服务侧的内容通过回声的方式写回
func echo(c net.Conn, input string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(input))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", input)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(input))

}
