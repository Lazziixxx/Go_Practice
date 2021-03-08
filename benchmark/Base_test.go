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
