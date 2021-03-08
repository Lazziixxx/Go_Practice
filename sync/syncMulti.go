package main

/*
关于lock和Unlock的使用
func Balance() int {
    mu.Lock()
    defer mu.Unlock()
    return balance
}
1:
使用defer  x.Unlock()保证锁的释放
同时这里充分说明了defer的特性
return balance 包含两个步骤  read balance + return value
真实顺序为：
mu.lock()
read balance to local
mu.Unlock()
return value of balance
2:即使临界区发生了panic ，defer Unlock依旧会执行
3：Go不支持可重入锁 Java支持 Why?
4：

func Withdraw(amount int) bool {
    mu.Lock()
    defer mu.Unlock()
    deposit(-amount)
    if balance < 0 {
        deposit(amount)
        return false // insufficient funds
    }
    return true
}

func Deposit(amount int) {
    mu.Lock()
    defer mu.Unlock()
    deposit(amount)
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }
如果在Withdraw中需要调用Deposit，可以实现对外不可见的函数deposit，避免了锁的重入
原则就是  在使用mutex时，需要保证临界区内的变量/函数/方法没有被导出
 */
