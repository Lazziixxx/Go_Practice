package main

import (
  "fmt"
  "testing"
)

/*************************不合适的关闭channel方法*************************/
func RudeSafeClose(ch chan int) (justClosed bool) {
  defer func() {
    if recover() != nil {
      fmt.Println("close a closed channel,need recover")
      justClosed = false //关闭一个已经closed的channel，justClosed 被改写成false return
    }
  }()
  
  close(ch)
  return true // justClosed = true; defer中justClosed = false; return
}

func testRudeSafeClose() {
  ch := make(chan int)
  close(ch)
  result := RudeSafeClose(ch) 
  fmt.Printf("close channel result :%t", result)
}

func TestRudeSafeClose(t *testing.T) {
  testRudeSafeClose()
}

/*
关于panic和recover
recover可以捕获panice传入的异常
func main() {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println(err) // 打印出badthing
    }
  }()
  f()
  ....
}

func f(){
  ....
  panic("badthing")
  ....
}
*/
/*************************不合适的关闭channel方法*************************/

/*************************sync.once关闭channel*************************/
/* sync.once对象 无论是否改变once.do(xxx)中的方法，都只会执行一次 */
type OnceChannel struct {
  task chan int
  once sync.once
}

func NewChannel() *OnceChannel{
  return &OnceChannel{task: make(chan int)}
}

func (mc *OnceChannel) SafeClose() {
  mc.once.do(func() {
    close(mc.task)
  })
}
/*************************sync.once关闭channel*************************/
/*************************sync.mutex关闭channel*************************/
type MutexChannel struct {
  task chan int
  closed bool
  mutex sync.Mutex
}

func IsClosed(ch *MutexChannel) bool {
  ch.mutex.Lock()
  defer ch.mutex.Unlock()
  return ch.closed
}

func (ch *MutexChannel) SafeClose() {
  ch.mutex.Lock()
  defer ch.mutex.Unlock()
  if !ch.closed {
    close(ch.task)
    ch.closed = true
  }
}
/*************************sync.mutex关闭channel*************************/

/**********************************************************************/
/************************纯通道操作安全的关闭通道************************/
/**********************************************************************/

/* 1:多接收者 单发送者 */
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  log.SetFlags(0)
  const Max = 100000
  const NumReceivers = 100
  wgReceivers := sync.WaitGroup{}
  wgReceivers.Add(NumReceivers)
  
  task := make(chan int)
  
  go func() {
    for {
      if value := rand.Intn(Max); value == 0 {
        close(task)
        return
      } else {
        task <- value
      }
    }
  }()
  
  for i := 0; i < NumReceivers; i++ {
    go func() {
      defer wgReceivers.Done()
      
      for value := range task {
        log.Println(value)
      }
      }()
    }
  
  wgReceivers.Wait()
}

/* 2:单接收者 多发送者 */
/* 接收者关闭通道，并使用一个多余的通道来通知接收者 */
package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  log.SetFlags(0)
  const Max = 100000
  const NumSenders = 100
  wgReceivers := sync.WaitGroup{}
  wgReceivers.Add(1)
  
  task := make(chan int)
  stopCh := make(chan struct{})
  
  for i := 0; i < NumSenders; i++ {
    go func() {
      for {
        /* 第一select块是保护作用 -- 因为即使stopCh被关闭了 第二个select的第一个分支有可能在若干个循环中不会被选中 */
        select {
          case <- stopCh:
            return
          default:      
        }

        select {
          case <- stopCh:
            return
          default:
          task <- rand.Intn(Max)
        }
      }
    }()
  }
  
  go func() {
    defer wgReceivers.Done()

    for value := range task {
      if value == Max - 1 {
        close(stopCh)
        return
      }
      log.Println(value)
    }
  }()
  
  wgReceivers.Wait()
}

/* 这种情况 没有关闭task任务通道；无需关闭，当一个通道不再被任何协程使用，会被回收掉 */

/* 3:多个接收者与发送者 */
/* tip：无缓冲队列：如果接收or发送一方没有ready，会堵塞双方 */
      /*有缓冲队列，如果缓冲区满，发送方堵塞，缓冲区空，接收放堵塞*/
package main

func main() {
  rand.Seed(time.Now().UnixNano())
  log.SetFlags(0)
  
  const Max = 1000
  const NumSenders = 10
  const NumReaders = 100
  
  wgReaders := sync.WaitGroup{}
  wgReaders.Add(NumReaders)
  
  /* 通道们 */
  task := make(chan int)
  stopCh := make(chan struct{})
  middler := make(chan string, 1) // 必须是有缓冲的通道
  var stopper string
  
  go func() {
    stopper = <- middler
    close(stopCh)
  }()
  
  /* 发送者的退出条件 */
  /* 1：通过stopCh感知到通道需要关闭 */
  /* 1：满足停止发送的条件 */
  for i := 0; i < NumSenders; i++ {
    go func(id string) {
      for {
        if value := rand.Intn(Max); value == 0 {
          select {
            case middler <- "Sender:" + id:
            default: // 如果已经有一个终止信号，其他发送者直接结束协程
          }
          return
        }
        
        select {
          case <- stopCh:
            return
          default:
        }
        
        select {
          case <- stopCh:
            return
          default:
            task <- rand.Intn(Max)
        }        
      }
    }(strconv.Itoa(i))  
  }

  /* 接收者的退出条件 */
  /* 1：通过stopCh感知到通道需要关闭 */
  /* 1：满足停止接收的条件 */
  for i := 0; i < NumReaders; i++ {
    go func(id string) {
      defer wgReaders.Done()
      for {
        select {
          case <- stopCh:
            return
          default:
        }
        
        select {
          case <- stopCh:
            return
          case value := <- task:
            if value == Max - 1 {
              select {
                case middler <- "Reader:" + id:
                default:
              }
              return
            }
        }
      }
    }(strconv.Itoa(i))
  }
  
  wgReaders.Wait()
  log.Println(stopper + "stop task")
}
