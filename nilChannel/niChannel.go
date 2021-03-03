/*
nil Channel的使用示例
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var testCh chan int  //当channel未分配内存时 就是nil channel
	fmt.Println(testCh)   //输出nil
	nilChan := make(chan int)
	go add(nilChan)
	go send(nilChan)
	time.Sleep(3 * time.Second)
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func add(c chan int) {
	sum := 0

	t := time.NewTimer(1 * time.Second)
	/*
	单次定时器的使用
	type Timer struct {
	    C <-chan Time     // The channel on which the time is delivered.
	    r runtimeTimer
	}
	在1s后 会向通过C中写入一个时间
	通过select语句对channel的监听，进行定时处理
	这里使用time.After也可以 与NewTimer(d).C等价
	t := time.After(1 * time.Second)
	for {
		select {
		....
		case <-t:
			c = nil
			......
		....
		}
	}
	 */
	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil //在定时时间后 将通过c置nil 关闭channel sum不会继续增加
			fmt.Println(sum)
	    }
	}
}


