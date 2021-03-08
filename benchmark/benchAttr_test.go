package main

import (
	"math/rand"
	"testing"
	"time"
)

func generate_c(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate_c(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b)}

/*
BenchmarkGenerate1000-12           47584             25046 ns/op           16376 B/op         11 allocs/op
BenchmarkGenerate10000-12           5721            206649 ns/op          386297 B/op         20 allocs/op
BenchmarkGenerate100000-12           600           2010018 ns/op         4654347 B/op         30 allocs/op
BenchmarkGenerate1000000-12           51          20464031 ns/op        45188456 B/op         44 allocs/op
根据输入的不同，输入增大10倍，每次函数调用的时长也差不多10倍，说明复杂度是线性的
*/

func fib_c(n int) int{
	if n == 1 || n == 0 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}

func BenchmarkFibNoSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib_c(30)
	}
}

func BenchmarkFibSleep(b *testing.B) {
	time.Sleep(time.Second * 3)//模拟某个耗时的任务
	for n := 0; n < b.N; n++ {
		fib_c(30)
	}
}

/*
BenchmarkFibSleep执行开销比BenchmarkFibNoSleep高很多 
中间增加了Sleep 模拟了其他耗时任务，会影响下面fib_c的性能统计
需要在定时器重置
*/

func BenchmarkFibResetTimer(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer() // 重置定时器
	for n := 0; n < b.N; n++ {
		fib_c(30) // run fib(30) b.N times
	}
}
