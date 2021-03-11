package main

import (
  "fmt"
  "runtime"
  "testing"
  "time"
)

func do(task chan int) {
  for {
    select {
    case <- task:
      time.Sleep(time.Millisecond)
      fmt.Println("task done")
    }
  }
}

func sendTask() {
  task := make(chan int)
  go do(task)
  for i := 0; i < 1000; i++ {
    task <- i
  }
}

func TestDo(t *testing.T) {
  t.Log(runtime.NumGoroutine())
  sendTask()
  time.Sleep(time.Second * 2)
  t.Log(runtime.NumGoroutine())
}

/*
 go test -run="TestDo" -v
=== RUN   TestDo
    goroutineExit_test.go:28: 2
    goroutineExit_test.go:31: 3
--- PASS: TestDo (3.52s)
执行结束后子协程多了一个，就是do协程，因为for中一直在等待新的task，然后sendtask已经结束了。

v, beforeClosed := <- ch
beforeClosed代表v是否是信道关闭前发送，是则为true，false代表信道已关闭
如果信道已关闭，<-ch永远不会发生堵塞。
*/

func doCloseTask(task chan int) {
  for {
    select {
      case _, chanIsClosed := <- task:
      if !chanIsClosed {
        fmt.Println("channel is closed")
        return
      }
      time.Sleep(time.Millisecond) 
    }
  }
}

func sendTaskWithClose() {
  task := make(chan int) 
  go doCloseTask(task)
  for i := 0; i < 1000; i++ {
    task <- i
  }
  close(task)
}

func TestDoClosed(t *testing.T) {
  t.Log(runtime.NumGoroutine())
  sendTaskWithClose()
  time.Sleep(time.Second * 2)
  t.Log(runtime.NumGoroutine())
}

/*
go test -run="TestDoClosed" -v
=== RUN   TestDoClosed
    goroutineExit_test.go:58: 2
channel is closed
    goroutineExit_test.go:61: 2
--- PASS: TestDoClosed (3.58s)
PASS
发送侧主动关闭channel，接收侧每次检测通道是否被关闭了，关闭了return退出协程；

=======通道和协程的垃圾回收========
注意，一个通道被其发送数据协程队列和接收数据协程队列中的所有协程引用着。因此，如果一个通道的这两个队列只要有一个不为空，则此通道肯定不会被垃圾回收。
另一方面，如果一个协程处于一个通道的某个协程队列之中，则此协程也肯定不会被垃圾回收，即使此通道仅被此协程所引用。事实上，一个协程只有在退出后才能被垃圾回收
*/
