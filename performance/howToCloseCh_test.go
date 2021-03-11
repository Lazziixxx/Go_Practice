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
