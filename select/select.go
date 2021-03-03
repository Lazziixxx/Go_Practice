package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutCheck()
	fullCheck()
	timerTool()
}

func timeoutCheck() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	select {
		case <-timeout:
			fmt.Println("timeout!")
	}
}

func fullCheck() {
	full := make(chan int, 1)
	full <- 1
	select {
		case full<-2:
	    default:
	    	fmt.Println("chan is Full.")
	}
}

func timerTool() {
	errChan := make(chan int)
	//ticker := time.NewTicker(2 * time.Second)
	ticker := time.Tick(2 * time.Second)
	go func() {
		time.Sleep(5 * time.Second)
		errChan <- 1
	}()
	LOOP:
	for {
		select {
		    //case <-ticker.C:
			case <- ticker:
				/* 定时间隔是2s */
				fmt.Println("Task is still running.")
			case err, ok := <-errChan:
				if ok {
					fmt.Printf("Task End: %d.", err)
					break LOOP
				}
		}
	}
}

/*
注意点1：
当For 和 select一起使用，break只能跳出select 无法跳出for
因此，如果没有break LOOP循环，会一直打印Task is still running
解决方法：使用标签，break + 标签或者goto + 标签

注意点2：
单独在select中不能使用continue 会编译错误
必须与for 配合使用

return 跟普通return一样，退出函数

注意点3：
这里time.Tick 和 time.NewTicker使用方法等价
 */
