package main

import "fmt"

func main() {
	counters := make(chan int)
	squares := make(chan int)
	go counter(counters) //隐式转换 将双方向channel转换成单向；
	go square(squares, counters) //不存在方向转换的语法，比如 不能将chan<- int 转换成chan int类型的双向channel
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i< 10; i++ {
		out <- i
	}
	close(out)
}

/*
单方向的channel类型
chan<- int 只用于发送int的channel
<-chan int只用于接受int的channel
限制条件将在编译的时候检测
只用发送者会关闭发送channel，close只接受的channel会产生编译错误
 */
func square(out chan<-int, in <-chan int) {
	for i:= range in {
		out<- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
