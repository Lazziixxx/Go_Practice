package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.RWMutex

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(time.Millisecond))
		fmt.Printf("Lock\n")
		m.Lock()
		fmt.Printf("UnLock\n")
		m.Unlock()
		ch<-1
	}()
	reader(4)
	<-ch
}

func reader(n int)  int{
	if n < 1 {
		return 0
	}
	time.Sleep(time.Duration(time.Millisecond))
	fmt.Printf("RdLock\n")
	m.RLock()
	defer func(){
		fmt.Printf("UnRdLock\n")
		m.RUnlock()
	}()
	return reader(n - 1) + n
}

/*
读者的递归调用导致死锁
RdLock
Lock
RdLock
fatal error: all goroutines are asleep - deadlock!
1：1个读者获取到锁
2：写者去申请锁，需要等第一个读者释放锁 才能申请到
3：又有一个读者去申请锁，需要等写者释放才能申请到，导致第一个读者无法释放；---- 第一个读者的释放依赖了写者的释放
造成deadlock
 */
