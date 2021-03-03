/* 简单多channel的使用 */
package main

import "fmt"

func main() {
	counters := make(chan int)
	squares := make(chan int)

	fmt.Println(counters)

	go func() {
		for i := 0; i < 5; i++ {
			counters <- i
		}
		close(counters)
	}()

	go func() {
		for  x:= range counters{
			squares <- x * x
		}
		close(squares)
	}()

	for {
		for x:= range squares {
			fmt.Println(x)
		}
	}

	return
}

/*
如果不主动close chan会导致goroutine死循环死锁
 */

/*
只有需要告诉接收者，所有数据全部发送结束 才需要主动关闭channel
channel在不被引用的时候会被GC回收

而文件打开操作，对于每个打开的文件，不使用的时候都要通过close关闭
 */

/*
重复关闭channel会导致panic ；panic: close of closed channel
关闭一个nil的channel也会导致panic； 关于何种情况channel为nil及nii channel的灵活使用 见nilChannel.go
 */

