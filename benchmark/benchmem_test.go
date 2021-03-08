/*
-benchmem可以用于度量内存分配次数
*/

package main

import(
  "math/rand"
  "testing"
  "time"
)

func generateWithCap(n int) []int {
  rand.Seed(time.Now().UnixNano())
  nums := make([]int, 0, n)
  for i := 0; i < n; i++ {
    nums = append(nums, rand.Int())
  }
}

func generate(n int) []int{
  rand.Seed(time.Now().UnixNano())
  nums := make([]int, 0)
  for i := 0; i < n; i++ {
    nums = append(nums, rand.Int())
  }
}

func BenchmarkGenerateWithCap(b *testing.B) {
  for n := 0; n < b.N; n++ {
    generateWithCap(1000000)
  }
}

func BenchmarkGenerate(b *testing.B) {
  for n := 0; n < b.N; n++ {
    generate(1000000)
  }
}

/*
go test -bench='Generate' .
goos: windows
goarch: amd64
BenchmarkGenerateWithCap-12           74          15848514 ns/op
BenchmarkGenerate-12                  56          20213343 ns/op
PASS
ok      _/D_/01_Ability/06_Go/TestProject/benchmark     2.389s

生成100w个随机数列 BenchmarkGenerateWithCap的性能高20%

增加-benchmem选项
go test -bench='Generate' -benchmem .
goos: windows
goarch: amd64
BenchmarkGenerateWithCap-12           74          15635555 ns/op         8003669 B/op          1 allocs/op
BenchmarkGenerate-12                  57          20410209 ns/op        45188441 B/op         42 allocs/op
PASS
ok      _/D_/01_Ability/06_Go/TestProject/benchmark     2.401s

BenchmarkGenerate内存分配了42次 分配的内存大小是BenchmarkGenerateWithCap的6倍
*/
