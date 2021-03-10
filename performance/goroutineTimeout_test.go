package main

func doBadthing(ch chan bool) {
  time.Sleep(time.Second)
  ch <- true
}

func timeout(f func(chan bool)) error {
  done := make(chan bool)
  go f(done)
  select {
    case <- done:
    fmt.Println("done")
    return nil
    case <- time.After(time.Millisecond): //特定延时，doBadthing中延时1s才会向done发送信号，此时没有接收者，发送者会一直堵塞，拉起的1000多个子协程无法退出，会导致内存耗尽；
    return fmt.Errorf("timeout")          //每个协程会占用2K内存
  }
}

func test(t *testing.T, f func(ch chan bool)) {
  t.Helper()
  for i := 0; i < 1000; i++ {
    timeout(f)
  }
  time.Sleep(time.Second * 2)
  t.Log(runtime.NumGoroutine())
}

func TestBadTimeout(t *testing.T) { test(t, doBadthing) }

/*
go test -v -run=TestBadTimeout
=== RUN   TestBadTimeout
    goroutineTimeout_test.go:36: 1002
--- PASS: TestBadTimeout (3.62s)

1002个子协程不能释放
*/

/*
如何避免协程超时无法退出？
*/
/* 1： 采用有buffer的channel */
func timeoutWithBuffer(f func(chan bool)) error {
  done := make(chan bool, 1) //创建有buffer的channel，即使没有接收者，发送者也不会被堵塞，不会导致协程无法退出
  go f(done)
  select {
    case <- done:
    fmt.Println("done")
    return nil
    case <- time.After(time.Millisecond): 
    fmt.Println("timeout")                
    return nil
  }
}

func testWithBuffer(t *testing.T, f func(ch chan bool)) {
  t.Helper()
  for i := 0; i < 1000; i++ {
    timeoutWithBuffer(f)
  }
  time.Sleep(time.Second * 2)
  t.Log(runtime.NumGoroutine())
}

func TestBadTimeoutWithBuffer(t *testing.T) { testWithBuffer(t, doBadthing) }

/*
    goroutineTimeout_test.go:60: 2
*/

/* 2: 采用select方式去写channel */
/*
从接收者 发送者的角度去考虑问题，有buffer的channel避免了接收者对发送者的堵塞
从另一个角度，如果发送者能够感知到是否有接收者再去决定是否发送，也可以解决协程无法退出的问题
*/
/* 采用select字段进行channel的写操作 无法写则直接退出 */
func doGoodthing(ch chan bool) {
  time.Sleep(time.Second)
  select {
    case ch <- true: //如果有接收者 才能写
    default:
      return
  }
}

func TestGoodTimeout(t *testing.T) { test(t, doGoodthing) }

/*
go test -v -run=TestGoodTimeout
=== RUN   TestGoodTimeout
    goroutineTimeout_test.go:71: 2
--- PASS: TestGoodTimeout (3.58s)
PASS
*/

/*
更为复杂且常用的场景，客户端发起请求，服务端将接收到的请求分为2段，一段执行任务，一段发送结果
*/

func doTwoWorks(work1, work2 chan bool) {
  time.Sleep(time.Second)
  select {
    case work1 <- true:
    default:
      return
  }
  time.Sleep(time.Second)
  work2 <- true
}

func timeOutWorkClient(f func(work1, work2 chan bool)) error {
  work1 := make(chan bool)
  work2 := make(chan bool)
  go f(work1, work2)
  select {
    case <- work1:
      <- work2
    fmt.Println("done")
    return nil
    case <-time.After(time.Millisecond):
    return fmt.Errorf("timeout")
  }
}

func TestTwoWorksTimeout(t *testing.T) {
  t.Helper()
  for i := 0; i < 1000; i++ {
    timeOutWorkClient(doTwoWorks)
  }
  time.Sleep(time.Second * 3)
  t.Log(runtime.NumGoroutine())
}

/*
1: work1正常执行，并向客户端返回结果
2：work1超时，并返回超时； 这里不能使用带buffer的channel，否则work1总能成功写入，work2总能执行到；
补充：缓冲区不能判断是否超时，select可以，写失败就说明没有接收者，就超时了；
如果需要用channel来实现客户端 服务端的同步，尽量不要使用带buffer的channel，其本质上是个异步操作；
*/

/*
总结：
goroutine不能被外部杀死，只能通过channel与他通信。A goroutine cannot be programmatically killed. It can only commit a cooperative suicide
建议：
1：尽量使用非阻塞 I/O（非阻塞 I/O 常用来实现高性能的网络库），阻塞 I/O 很可能导致 goroutine 在某个调用一直等待，而无法正确结束。
2：业务逻辑总是考虑退出机制，避免死循环。
3：任务分段执行，超时后即时退出，避免 goroutine 无用的执行过多，浪费资源
*/
