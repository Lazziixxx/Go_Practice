package main

import "testing"

func fib(n int) int {
  if n == 0 || n == 1 {
    return n
  }
  return fib(n-2) + fib(n-1)
}

func BenchmarkFib(b *testing.B) {
  for n:=0; n < b.N; n++ {
    fib(30)
  }
}

/*
go test -bench . 
运行当前目录下的benchmark用例
go test -bench='xxx$' .
运行以xxx结尾的benchmark用例
修改绑核：
go test -bench='xxx$' -cpu=2,4 .
增加-cpu选型 即GOMAXPROCS CPU核数
提升测试准确度:
go test -bench='xxx$' -benchtime=5s .
benchmark默认测试时间是1s 可以使用-benchtime=5s指定测试时间，提升准确度（实际测试时间比5s会长，编译执行销毁需要时间）
-benchtime=100x  也可以指定为测试100次
go test -bench='xxx$' -benchtime=5s -count=3 .
设置进行benchmark的轮数 进行三伦benchmark
*/
