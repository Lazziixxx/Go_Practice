package main

import (
  "math/rand"
  "testing"
  "time"
)

func generateWithCap(n int) []int {
  rand.Seed(time.Now().UnixNano)
  nums := make([]int, 0, n)
  for i := 0; i < n ; i++ {
    nums := append(nums, rand.Int())
  }
  return nums
}

func BenchmarkSortNums(b *testing.B) {
  for n := 0; n < b.N; n++ {
    b.StopTimer()
    nums := generateWithCap(10000)
    b.StartTimer()
    sort.Ints(nums)
  }
}
