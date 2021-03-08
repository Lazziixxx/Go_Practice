package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

/*
打印出同一时间出现的读者和写者
 */

func main() {
	var m sync.RWMutex
	readCh := make(chan int)
	writeCh := make(chan int)
	var rn,wn int
	go func() {
		for {
			select {
				case r := <-readCh:
					rn += r
				case w := <-writeCh:
					wn += w
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", rn), strings.Repeat("W", wn))//打印当前时刻存在读者和写者数目
		}
	}()

	var w sync.WaitGroup
	/*
	计数器 主goroutine用于等待所有的子goroutine结束
	每生成一个子协程:add()
	子协程结束：done()
	主协程：wait()
	 */
	for i:=0; i < 10; i++ {
		w.Add(1)
		go Reader(readCh, &m,&w)
	}

	for i:=0; i < 3; i++ {
		w.Add(1)
		go Writer(writeCh, &m,&w)
	}
	w.Wait()
}

func Reader(ch chan  int, m *sync.RWMutex,w *sync.WaitGroup) {
	m.RLock()
	defer m.RUnlock()
	ch <- 1
	sleep()
	ch <- -1
	w.Done()
}

func Writer(ch chan  int, m *sync.RWMutex,w *sync.WaitGroup) {
	m.Lock()
	defer m.Unlock()
	ch <- 1
	sleep()
	ch <- -1
	w.Done()
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func init() {
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
}

/*
R
RR
R

W

R
RR
RRR
RRRR
RRRRR
RRRRRR
RRRRRRR
RRRRRRRR
RRRRRRR
RRRRRR
RRRRR
RRRR
RRR
RR
R

W

W

输出结果理解：
1：select语句----当临界区内的go routinue数目发生变化，程序就会换行
2：每行的显示说明允许多个读者访问或者一个写者
3：写者调用Lock后，后续的读者会被阻塞，临界区内的读者会依次结束离开，再锁定临界区 --- 保证了写者不会饿死
4：当写者结束后，阻塞的读者会先获得临界区的访问权 --- 保证了读者不会被饿死
 */
